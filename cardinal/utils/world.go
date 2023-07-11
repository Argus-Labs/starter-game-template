package utils

import (
	"os"

	ecs "github.com/argus-labs/world-engine/cardinal/ecs"
	"github.com/argus-labs/world-engine/cardinal/ecs/inmem"
	"github.com/argus-labs/world-engine/cardinal/ecs/storage"
)

var (
	EnvRedisAddr = os.Getenv("REDIS_ADDR")
)

// newDevWorld is the recommended way of running the game for development
// where you are going to need use Retool to inspect the state.
// NOTE(1): You will need to have Redis running in `EnvRedisAddr` for this to work.
// NOTE(2): In prod, your Redis should have a password loaded from env var so don't use this.
func NewDevWorld() *ecs.World {
	rs := storage.NewRedisStorage(storage.Options{
		Addr:     EnvRedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	}, "in-memory-world")
	worldStorage := storage.NewWorldStorage(&rs)
	world, err := ecs.NewWorld(worldStorage)
	if err != nil {
		panic(err)
	}

	return world
}

// newInmemWorld is the most convenient way to run the game locally
// because it doesn't require spinning up Redis in a container
// it runs a Redis server as a part of the Go process
// However, it will not work with Retool.
func NewInmemWorld() *ecs.World {
	return inmem.NewECSWorld()
}
