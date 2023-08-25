package component

import "pkg.world.dev/world-engine/cardinal/ecs"

type PlayerComponent struct {
	Nickname string
}

var Player = ecs.NewComponentType[PlayerComponent]()
