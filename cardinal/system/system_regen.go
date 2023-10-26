package system

import (
	comp "github.com/argus-labs/starter-game-template/cardinal/component"
	"pkg.world.dev/world-engine/cardinal"
)

// RegenSystem is a system that replenishes the player's HP at every tick.
// This provides a simple example of how to create a system that doesn't rely on a transaction to update a component.
func RegenSystem(wCtx cardinal.WorldContext) error {
	q, err := wCtx.NewSearch(cardinal.Exact(comp.PlayerComponent{}, comp.HealthComponent{}))
	if err != nil {
		return err
	}
	err = q.Each(wCtx, func(id cardinal.EntityID) bool {
		// Get the health component for the player
		health, err := cardinal.GetComponent[comp.HealthComponent](wCtx, id)
		if err != nil {
			return true
		}

		// Replenish some HP and update the component
		health.HP += 1
		if err := cardinal.SetComponent[comp.HealthComponent](wCtx, id, health); err != nil {
			return true
		}

		return true
	})
	if err != nil {
		return err
	}

	return nil
}
