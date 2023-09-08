package game

import (
	"github.com/argus-labs/starter-game-template/cardinal/component"
	"pkg.world.dev/world-engine/cardinal"
)

type IArchetype struct {
	Label      string
	Components []cardinal.AnyComponentType
}

var (
	Archetypes = []IArchetype{playerArchetype}

	playerArchetype = IArchetype{
		Label:      "player",
		Components: []cardinal.AnyComponentType{component.Player, component.Health},
	}
)
