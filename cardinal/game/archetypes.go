package game

import (
	"github.com/argus-labs/starter-game-template/component"
	"github.com/argus-labs/starter-game-template/types"
	ecs "github.com/argus-labs/world-engine/cardinal/ecs"
)

// This is where we define the archetypes for our game.
// The primary use of this is primarily to create an archetype query
// that makes it easier for you to query for entities that have a certain set of components
// through the use of the `query_archetype` message.

var (
	Archetypes = []types.IArchetype{playerArchetype}

	playerArchetype = types.IArchetype{
		Label:      "player",
		Components: []ecs.IComponentType{component.Player, component.Health},
	}
)
