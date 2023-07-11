package component

import "github.com/argus-labs/world-engine/cardinal/ecs"

type PlayerComponent struct {
	Tag string
}

var Player = ecs.NewComponentType[PlayerComponent]()
