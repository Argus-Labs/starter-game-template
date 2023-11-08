package system

import (
	"fmt"

	comp "github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/tx"
	"pkg.world.dev/world-engine/cardinal"
)

// PlayerSpawnerSystem is a system that spawns players based on `CreatePlayer` transactions.
// This provides a simple example of how to create a system that creates a new entity.
func PlayerSpawnerSystem(world cardinal.WorldContext) error {
	// Get all the transactions that are of type CreatePlayer from the tx queue
	createTxs := tx.CreatePlayer.In(world)

	// Iterate through all transactions and process them individually.
	// DEV: it's important here that you don't break out of the loop or return an error here
	// or otherwise the rest of the transaction will not be processed & get dropped.
	// In the future, you will be able to add error receipts to transaction receipts.
	for i, create := range createTxs {
		id, err := cardinal.Create(world, comp.Player{}, comp.Health{})
		if err != nil {
			tx.CreatePlayer.AddError(world, create.Hash(),
				fmt.Errorf("error creating player: %w", err))
			continue
		}
		err = cardinal.SetComponent[comp.Player](world, id, &comp.Player{Nickname: create.Value().Nickname})
		if err != nil {
			tx.CreatePlayer.AddError(world, create.Hash(),
				fmt.Errorf("error setting player nickname: %w", err))
			continue
		}

		err = cardinal.SetComponent[comp.Health](world, id, &comp.Health{HP: 100})
		if err != nil {
			tx.CreatePlayer.AddError(world, create.Hash(),
				fmt.Errorf("error setting player health: %w", err))
			continue
		}
		tx.CreatePlayer.SetResult(world, create.Hash(), tx.CreatePlayerMsgReply{Success: true})
		world.EmitEvent(fmt.Sprintf("%d player: %d created, %d/%d", i+1, id, i+1, len(createTxs)))
	}

	return nil
}
