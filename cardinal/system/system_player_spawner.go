package system

import (
	"fmt"

	comp "github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/msg"
	"pkg.world.dev/world-engine/cardinal"
)

// PlayerSpawnerSystem spawns players based on `CreatePlayer` transactions.
// This provides an example of a system that creates a new entity.
func PlayerSpawnerSystem(world cardinal.WorldContext) error {
	msg.CreatePlayer.Each(world, func(create cardinal.TxData[msg.CreatePlayerMsg]) (msg.CreatePlayerResult, error) {
		id, err := cardinal.Create(world, comp.Player{}, comp.Health{})
		if err != nil {
			return msg.CreatePlayerResult{}, fmt.Errorf("error creating player: %w", err)
		}

		err = cardinal.SetComponent[comp.Player](world, id, &comp.Player{Nickname: create.Msg().Nickname})
		if err != nil {
			return msg.CreatePlayerResult{}, fmt.Errorf("error setting player nickname: %w", err)
		}

		err = cardinal.SetComponent[comp.Health](world, id, &comp.Health{HP: 100})
		if err != nil {
			return msg.CreatePlayerResult{}, fmt.Errorf("error setting player health: %w", err)
		}

		world.EmitEvent(fmt.Sprintf("new player %d created", id))
		return msg.CreatePlayerResult{Success: true}, nil
	})
	return nil
}
