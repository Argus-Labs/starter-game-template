package system

import (
	"fmt"

	comp "github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/tx"
	"pkg.world.dev/world-engine/cardinal"
)

// AttackSystem is a system that inflict damage to player's HP based on `AttackPlayer` transactions.
// This provides a simple example of how to create a system that modifies the component of an entity.
func AttackSystem(world cardinal.WorldContext) error {
	// Get all the transactions that are of type CreatePlayer from the tx queue

	// Create an index of player tags to its health component
	playerTagToID := map[string]cardinal.EntityID{}
	q, err := world.NewSearch(cardinal.Exact(comp.Player{}, comp.Health{}))
	if err != nil {
		return err
	}
	err = q.Each(world, func(id cardinal.EntityID) bool {
		player, err := cardinal.GetComponent[comp.Player](world, id)
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
	tx.AttackPlayer.ForEach(world, func(attack cardinal.TxData[tx.AttackPlayerMsg]) (tx.AttackPlayerMsgReply, error) {
		target := attack.Value().TargetNickname
		targetPlayerID, ok := playerTagToID[target]
		// If the target player doesn't exist, skip this transaction
		if !ok {
			return tx.AttackPlayerMsgReply{}, fmt.Errorf("target %q does not exist", target)
		}

		// Get the health component for the target player
		health, err := cardinal.GetComponent[comp.Health](world, targetPlayerID)
		if err != nil {
			return tx.AttackPlayerMsgReply{}, fmt.Errorf("can't get health for %q: %w", target, err)
		}

		// Inflict damage and update the component
		health.HP -= 10
		if err := cardinal.SetComponent[comp.Health](world, targetPlayerID, health); err != nil {
			return tx.AttackPlayerMsgReply{}, fmt.Errorf("failed to set health on %q: %w", target, err)
		}
		return tx.AttackPlayerMsgReply{}, nil
	})

	return nil
}
