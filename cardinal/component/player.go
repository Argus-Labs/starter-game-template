package component

import "github.com/argus-labs/world-engine/cardinal/ecs"

type PlayerComponent struct {
	Nickname string
}

var Player = ecs.NewComponentType[PlayerComponent]()
