package component

import "pkg.world.dev/world-engine/cardinal/ecs"

type HealthComponent struct {
	HP int
}

var Health = ecs.NewComponentType[HealthComponent]("Health")
