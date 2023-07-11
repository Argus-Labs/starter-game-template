package main

import (
	"github.com/argus-labs/starter-game-template/component"
	"github.com/argus-labs/starter-game-template/msg"
	tx "github.com/argus-labs/starter-game-template/msg/tx"
	"github.com/argus-labs/starter-game-template/system"
	"github.com/argus-labs/starter-game-template/utils"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// NOTE: Uses a Redis container
	// Best to use this for testing with Retool
	world := utils.NewDevWorld()

	// NOTE: If you want to use an in-memory Redis, use this instead.
	// This is the easiest way to run Cardinal locally, but doen't work with Retool.
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
		tx.TxCreatePlayer,
		tx.TxAttackPlayer,
	))

	// Each system executes deterministically in the order they are added.
	// This is a neat feature that can be straegically used for systems that depends on the order of execution.
	// For example, you may want to run the attack system before the regen system
	// so that the player's HP is subtracted (and player killed if it reaches 0) before HP is regenerated.
	world.AddSystem(system.AttackSystem)
	world.AddSystem(system.RegenSystem)
	world.AddSystem(system.PlayerSpawnerSystem)

	// Load game state
	utils.Must(world.LoadGameState())

	// Start game loop as a goroutine
	go utils.GameLoop(world)

	// Register message handlers
	h := msg.NewMsgHandler(world)
	utils.RegisterRpc(utils.GetPort(), func() utils.CardinalHandlers {
		return utils.CardinalHandlers{
			{"query_archetype", h.Query.Archetype},
			{"query_constant", h.Query.Constant},
			{"tx_create_player", h.Tx.CreatePlayer},
			{"tx_attack_player", h.Tx.AttackPlayer},
		}
	}())
}
