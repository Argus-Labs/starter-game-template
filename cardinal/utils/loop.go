package utils

import (
	"github.com/argus-labs/world-engine/cardinal/ecs"
	"log"
	"time"
)

// TODO: this should probably be upstreamed to the ecs lib
func GameLoop(world *ecs.World) {
	log.Print("Starting game loop...")
	for range time.Tick(time.Second) {
		if err := world.Tick(); err != nil {
			panic(err)
		}
	}
}
