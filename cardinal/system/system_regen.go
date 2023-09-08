package system

import (
	comp "github.com/argus-labs/starter-game-template/cardinal/component"
	"pkg.world.dev/world-engine/cardinal"
)

// RegenSystem is a system that replenishes the player's HP at every tick.
// This provides a simple example of how to create a system that doesn't rely on a transaction to update a component.
func RegenSystem(world *cardinal.World, _ *cardinal.TransactionQueue, _ *cardinal.Logger) error {
	cardinal.NewQuery(cardinal.Exact(comp.Player, comp.Health)).Each(world, func(id cardinal.EntityID) bool {
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
