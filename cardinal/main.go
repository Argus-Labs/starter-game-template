package main

import (
	"context"
	"github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/read"
	"github.com/argus-labs/starter-game-template/cardinal/system"
	"github.com/argus-labs/starter-game-template/cardinal/tx"
	"github.com/argus-labs/starter-game-template/cardinal/utils"
	"github.com/argus-labs/world-engine/cardinal/server"
	"github.com/rs/zerolog"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cfg := GetConfig()

	// TODO: In production, you should set DEPLOY_MODE=production
	// and set REDIS_ADDR and REDIS_PASSWORD to use a real Redis instance.
	// Otherwise, by default cardinal will run using an in-memory "miniredis"
	world := cfg.World

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

	// Register read endpoints
	// NOTE: You must register your read endpoints here,
	// otherwise it will not be accessible.
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
	world.StartGameLoop(context.Background(), time.Second)

	// TODO: When launching to production, you should enable signature verification.
	h, err := server.NewHandler(world, server.DisableSignatureVerification())
	if err != nil {
		panic(err)
	}
	h.Serve("", cfg.CardinalPort)
}
