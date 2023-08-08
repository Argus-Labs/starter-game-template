package system

import (
	"fmt"
	comp "github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/tx"
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

// PlayerSpawnerSystem is a system that spawns players based on `CreatePlayer` transactions.
// This provides a simple example of how to create a system that creates a new entity.
func PlayerSpawnerSystem(world *ecs.World, tq *ecs.TransactionQueue) error {
	// Get all the transactions that are of type CreatePlayer from the tx queue
	createTxs := tx.CreatePlayer.In(tq)

	// Iterate through all transactions and process them individually.
	// DEV: it's important here that you don't break out of the loop or return an error here
	// or otherwise the rest of the transaction will not be processed & get dropped.
	// In the future, you will be able to add error receipts to transaction receipts.
	for _, create := range createTxs {
		id, err := world.Create(comp.Player, comp.Health)
		if err != nil {
			fmt.Println("Error creating player")
			continue
		}

		err = comp.Player.Set(world, id, comp.PlayerComponent{Nickname: create.Nickname})
		if err != nil {
			fmt.Println("Error setting player nickname")
			continue
		}

		err = comp.Health.Set(world, id, comp.HealthComponent{HP: 100})
		if err != nil {
			fmt.Println("Error setting player health")
			continue
		}
	}

	return nil
}
