package utils

import (
	"github.com/rs/zerolog/log"

	"pkg.world.dev/world-engine/cardinal/ecs"
	"pkg.world.dev/world-engine/cardinal/ecs/inmem"
	"pkg.world.dev/world-engine/cardinal/ecs/storage"
)

// NewWorld is the recommended way to run the game
func NewWorld(addr string, password string, deployMode string) *ecs.World {
	options := make([]ecs.Option, 0)
	log.Log().Msg("Running in normal mode, using external Redis")
	if addr == "" {
		log.Log().Msg("Redis address is not set, using fallback - localhost:6379")
		addr = "localhost:6379"
	}
	if password == "" {
		log.Log().Msg("Redis password is not set, make sure to set up redis with password in prod")
		password = ""
	}
	if deployMode == "development" {
		options = append(options, ecs.WithPrettyLog())
	}

	rs := storage.NewRedisStorage(storage.Options{
		Addr:     addr,
		Password: password, // make sure to set this in prod
		DB:       0,        // use default DB
	}, "world")
	worldStorage := storage.NewWorldStorage(&rs)
	world, err := ecs.NewWorld(worldStorage, options...)
	if err != nil {
		panic(err)
	}

	return world
}

// NewEmbeddedWorld is the most convenient way to run the game
// because it doesn't require spinning up Redis in a container.
// It runs a Redis server as a part of the Go process.
// NOTE: worlds with embedded redis are incompatible with Cardinal Editor.
func NewEmbeddedWorld(deployMode string) *ecs.World {
	log.Log().Msg("Running in embedded mode, using embedded miniredis")
	options := make([]ecs.Option, 0)
	if deployMode == "development" {
		options = append(options, ecs.WithPrettyLog())
	}
	return inmem.NewECSWorld(options...)
}
