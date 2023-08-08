package system

import (
	comp "github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/world-engine/cardinal/ecs"
	"github.com/argus-labs/world-engine/cardinal/ecs/filter"
	"github.com/argus-labs/world-engine/cardinal/ecs/storage"
)

// RegenSystem is a system that replenishes the player's HP at every tick.
// This provides a simple example of how to create a system that doesn't rely on a transaction to update a component.
func RegenSystem(world *ecs.World, _ *ecs.TransactionQueue) error {
	ecs.NewQuery(filter.Exact(comp.Player, comp.Health)).Each(world, func(id storage.EntityID) bool {
		// Get the health component for the player
		health, err := comp.Health.Get(world, id)
		if err != nil {
			return true
		}

		// Replenish some HP and update the component
		health.HP += 1
		if err := comp.Health.Set(world, id, health); err != nil {
			return true
		}

		return true
	})

	return nil
}
