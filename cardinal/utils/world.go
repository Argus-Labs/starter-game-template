package utils

import (
	"github.com/argus-labs/world-engine/cardinal/ecs"
	"github.com/argus-labs/world-engine/cardinal/ecs/inmem"
	"github.com/argus-labs/world-engine/cardinal/ecs/storage"
	"github.com/rs/zerolog/log"
)

// NewWorld is the recommended way to run the game
func NewWorld(addr string, password string) *ecs.World {
	log.Log().Msg("Running in normal mode, using external Redis")
	if addr == "" {
		log.Log().Msg("Redis address is not set, using fallback - localhost:6379")
		addr = "localhost:6379"
	}
	if password == "" {
		log.Log().Msg("Redis password is not set, make sure to set up redis with password in prod")
		password = ""
	}

	rs := storage.NewRedisStorage(storage.Options{
		Addr:     addr,
		Password: password, // make sure to set this in prod
		DB:       0,        // use default DB
	}, "world")
	worldStorage := storage.NewWorldStorage(&rs)
	world, err := ecs.NewWorld(worldStorage)
	if err != nil {
		panic(err)
	}

	return world
}

// NewEmbeddedWorld is the most convenient way to run the game
// because it doesn't require spinning up Redis in a container
// it runs a Redis server as a part of the Go process
// However, it will not work with Cardinal Editor.
func NewEmbeddedWorld() *ecs.World {
	log.Log().Msg("Running in embedded mode, using embedded miniredis")
	return inmem.NewECSWorld()
}
