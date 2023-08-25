package game

import (
	"github.com/argus-labs/starter-game-template/cardinal/component"
	"pkg.world.dev/world-engine/cardinal/ecs"
)

type IArchetype struct {
	Label      string
	Components []ecs.IComponentType
}

var (
	Archetypes = []IArchetype{playerArchetype}

	playerArchetype = IArchetype{
		Label:      "player",
		Components: []ecs.IComponentType{component.Player, component.Health},
	}
)
