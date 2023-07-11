package component

import "github.com/argus-labs/world-engine/cardinal/ecs"

type HealthComponent struct {
	HP int
}

var Health = ecs.NewComponentType[HealthComponent]()
