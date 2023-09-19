package component

import "pkg.world.dev/world-engine/cardinal/ecs"

type PlayerComponent struct {
	Nickname string `json:"nickname"`
}

var Player = ecs.NewComponentType[PlayerComponent]("Player")
