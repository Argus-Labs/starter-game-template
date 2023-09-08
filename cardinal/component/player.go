package component

import "pkg.world.dev/world-engine/cardinal"

type PlayerComponent struct {
	Nickname string
}

var Player = cardinal.NewComponentType[PlayerComponent]()
