package system

import (
	comp "github.com/argus-labs/starter-game-template/component"
	msg "github.com/argus-labs/starter-game-template/msg/tx"
	"github.com/argus-labs/world-engine/cardinal/ecs"
	"github.com/argus-labs/world-engine/cardinal/ecs/filter"
	"github.com/argus-labs/world-engine/cardinal/ecs/storage"
)

// AttackSystem is a system that inflict damage  to players's HP based on `TxAttackPlayer` transactions.
// This provides a simple example of how to create a system that modifies the component of an entity.
func AttackSystem(world *ecs.World, tq *ecs.TransactionQueue) error {
	// Get all the transactions that are of type TxCreatePlayer from the tx queue
	attackTxs := msg.TxAttackPlayer.In(tq)

	// Create an index of player tags to its health component
	playerTagToID := map[string]storage.EntityID{}
	ecs.NewQuery(filter.Exact(comp.Player, comp.Health)).Each(world, func(id storage.EntityID) {
		player, err := comp.Player.Get(world, id)
		if err != nil {
			return
		}

		playerTagToID[player.Tag] = id
	})

	// Iterate through all transactions and process them individually.
	// DEV: it's important here that you don't break out of the loop or return an error here
	// or otherwise the rest of the transaction will not be processed & get dropped.
	// In the future, you will be able to add error receipts to transaction receipts.
	for _, tx := range attackTxs {
		targetPlayerID, ok := playerTagToID[tx.TargetPlayerTag]
		// If the target player doesn't exist, skip this transaction
		if !ok {
			continue
		}

		// Get the health component for the target player
		health, err := comp.Health.Get(world, targetPlayerID)
		if err != nil {
			continue
		}

		// Inflict damage and update the component
		health.HP -= 10
		if err := comp.Health.Set(world, targetPlayerID, health); err != nil {
			continue
		}
	}

	return nil
}
