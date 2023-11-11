package main

import (
	"errors"
	"os"

	"github.com/argus-labs/starter-game-template/cardinal/query"

	"github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/system"
	"github.com/argus-labs/starter-game-template/cardinal/tx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"pkg.world.dev/world-engine/cardinal"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg := GetConfig()

	// TODO: In production, you should set DEPLOY_MODE=production
	// and set REDIS_ADDR and REDIS_PASSWORD to use a real Redis instance.
	// Otherwise, by default cardinal will run using an in-memory redis.
	// TODO: When launching to production, you should enable signature verification.
	w := NewWorld(cfg, cardinal.WithDisableSignatureVerification())

	// Register components
	// NOTE: You must register your components here,
	// otherwise it will show an error when you try to use them in a system.
	Must(
		cardinal.RegisterComponent[component.Player](w),
		cardinal.RegisterComponent[component.Health](w),
	)

	// Register transactions
	// NOTE: You must register your transactions here,
	// otherwise it will show an error when you try to use them in a system.
	Must(cardinal.RegisterTransactions(w,
		tx.CreatePlayer,
		tx.AttackPlayer,
	))

	// Register read endpoints
	// NOTE: You must register your read endpoints here,
	// otherwise it will not be accessible.
	Must(cardinal.RegisterQueries(w,
		query.Constant,
	))

	// Each system executes deterministically in the order they are added.
	// This is a neat feature that can be strategically used for systems that depends on the order of execution.
	// For example, you may want to run the attack system before the regen system
	// so that the player's HP is subtracted (and player killed if it reaches 0) before HP is regenerated.
	cardinal.RegisterSystems(w,
		system.AttackSystem,
		system.RegenSystem,
		system.PlayerSpawnerSystem,
	)

	Must(w.StartGame())
}

func Must(err ...error) {
	e := errors.Join(err...)
	if err != nil {
		log.Fatal().Err(e)
	}
}
