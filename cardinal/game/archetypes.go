package game

import (
	"github.com/argus-labs/starter-game-template/cardinal/component"
	"pkg.world.dev/world-engine/cardinal/ecs/component_metadata"
)

type IArchetype struct {
	Label      string
	Components []component_metadata.Component
}

var (
	Archetypes = []IArchetype{playerArchetype}

	playerArchetype = IArchetype{
		Label:      "player",
		Components: []component_metadata.Component{&component.PlayerComponent{}, &component.HealthComponent{}},
	}
)
