package main

import (
	"github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/read"
	"github.com/argus-labs/starter-game-template/cardinal/system"
	"github.com/argus-labs/starter-game-template/cardinal/tx"
	"github.com/argus-labs/starter-game-template/cardinal/utils"
	"github.com/argus-labs/world-engine/cardinal/server"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cfg := utils.GetConfig()

	// NOTE: Uses a Redis container
	// Best to use this for testing with Retool
	world := cfg.World

	// NOTE: If you want to use an in-memory Redis, use this instead.
	// This is the easiest way to run Cardinal locally, but does not work with Retool.
	// world := utils.NewInmemWorld()

	// Register components
	// NOTE: You must register your components here,
	// otherwise it will show an error when you try to use them in a system.
	utils.Must(world.RegisterComponents(
		component.Player,
		component.Health,
	))

	// Register transactions
	// NOTE: You must register your transactions here,
	// otherwise it will show an error when you try to use them in a system.
	utils.Must(world.RegisterTransactions(
		tx.CreatePlayer,
		tx.AttackPlayer,
	))

	utils.Must(world.RegisterReads(
		read.Archetype,
		read.Constant,
	))

	// Each system executes deterministically in the order they are added.
	// This is a neat feature that can be strategically used for systems that depends on the order of execution.
	// For example, you may want to run the attack system before the regen system
	// so that the player's HP is subtracted (and player killed if it reaches 0) before HP is regenerated.
	world.AddSystem(system.AttackSystem)
	world.AddSystem(system.RegenSystem)
	world.AddSystem(system.PlayerSpawnerSystem)

	// Load game state
	utils.Must(world.LoadGameState())

	// Start game loop as a goroutine
	go utils.GameLoop(world)

	// TODO: When launching to production, you should enable signature verification.
	h, err := server.NewHandler(world, server.DisableSignatureVerification())
	if err != nil {
		panic(err)
	}
	h.Serve("", cfg.CardinalPort)
}
