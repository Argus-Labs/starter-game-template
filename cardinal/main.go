package main

import (
	"errors"
	"fmt"

	"github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/read"
	"github.com/argus-labs/starter-game-template/cardinal/system"
	"github.com/argus-labs/starter-game-template/cardinal/tx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"pkg.world.dev/world-engine/cardinal"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cfg := GetConfig()

	// TODO: In production, you should set DEPLOY_MODE=production
	// and set REDIS_ADDR and REDIS_PASSWORD to use a real Redis instance.
	// Otherwise, by default cardinal will run using an in-memory redis.
	// TODO: When launching to production, you should enable signature verification.
	fmt.Println("Serving Cardinal at: ", cfg.CardinalPort)
	world := NewWorld(cfg, cardinal.WithDisableSignatureVerification())

	// Register components
	// NOTE: You must register your components here,
	// otherwise it will show an error when you try to use them in a system.
	err := errors.Join(
		cardinal.RegisterComponent[component.PlayerComponent](world),
		cardinal.RegisterComponent[component.HealthComponent](world))
	if err != nil {
		log.Fatal().Err(err)
	}

	// Register transactions
	// NOTE: You must register your transactions here,
	// otherwise it will show an error when you try to use them in a system.
	err = cardinal.RegisterTransactions(world,
		tx.CreatePlayer,
		tx.AttackPlayer,
	)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Register read endpoints
	// NOTE: You must register your read endpoints here,
	// otherwise it will not be accessible.
	err = cardinal.RegisterQueries(world,
		read.Archetype,
		read.Constant,
	)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Each system executes deterministically in the order they are added.
	// This is a neat feature that can be strategically used for systems that depends on the order of execution.
	// For example, you may want to run the attack system before the regen system
	// so that the player's HP is subtracted (and player killed if it reaches 0) before HP is regenerated.

	cardinal.RegisterSystems(world,
		system.AttackSystem,
		system.RegenSystem,
		system.PlayerSpawnerSystem)

	err = world.StartGame()
	if err != nil {
		panic(err)
	}

}
