package utils

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/argus-labs/world-engine/cardinal/ecs"
)

func GameLoop(world *ecs.World) {
	log.Info().Msg("Game loop started")
	for range time.Tick(time.Second) {
		if err := world.Tick(context.Background()); err != nil {
			panic(err)
		}
	}
}
