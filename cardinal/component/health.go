package component

import (
	"pkg.world.dev/world-engine/cardinal"
)

type HealthComponent struct {
	HP int
}

var Health = cardinal.NewComponentType[HealthComponent]()
