package system

import (
	"fmt"

	comp "github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/tx"
	"pkg.world.dev/world-engine/cardinal"
)

// AttackSystem is a system that inflict damage to player's HP based on `AttackPlayer` transactions.
// This provides a simple example of how to create a system that modifies the component of an entity.
func AttackSystem(wCtx cardinal.WorldContext) error {
	// Get all the transactions that are of type CreatePlayer from the tx queue
	attackTxs := tx.AttackPlayer.In(wCtx)

	// Create an index of player tags to its health component
	playerTagToID := map[string]cardinal.EntityID{}
	q, err := wCtx.NewSearch(cardinal.Exact(comp.PlayerComponent{}, comp.HealthComponent{}))
	if err != nil {
		return err
	}
	err = q.Each(wCtx, func(id cardinal.EntityID) bool {
		player, err := cardinal.GetComponent[comp.PlayerComponent](wCtx, id)
		if err != nil {
			return true
		}

		playerTagToID[player.Nickname] = id
		return true
	})
	if err != nil {
		return err
	}

	// Iterate through all transactions and process them individually.
	// DEV: it's important here that you don't break out of the loop or return an error here
	// or otherwise the rest of the transaction will not be processed & get dropped.
	// In the future, you will be able to add error receipts to transaction receipts.
	for _, attack := range attackTxs {
		target := attack.Value().TargetNickname
		targetPlayerID, ok := playerTagToID[target]
		// If the target player doesn't exist, skip this transaction
		if !ok {
			tx.AttackPlayer.AddError(wCtx, attack.Hash(),
				fmt.Errorf("target %q does not exist", target))
			continue
		}

		// Get the health component for the target player
		health, err := cardinal.GetComponent[comp.HealthComponent](wCtx, targetPlayerID)
		if err != nil {
			tx.AttackPlayer.AddError(wCtx, attack.Hash(),
				fmt.Errorf("can't get health for %q: %w", target, err))
			continue
		}

		// Inflict damage and update the component
		health.HP -= 10
		if err := cardinal.SetComponent[comp.HealthComponent](wCtx, targetPlayerID, health); err != nil {
			tx.AttackPlayer.AddError(wCtx, attack.Hash(),
				fmt.Errorf("failed to set health on %q: %w", target, err))
			continue
		}
	}

	return nil
}
