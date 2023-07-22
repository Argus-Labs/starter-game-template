package game

import (
	"github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/world-engine/cardinal/ecs"
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
