package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/message"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/msg"
)

const (
	InitialHP = 100
)

// PlayerSpawnerSystem spawns players based on `CreatePlayer` transactions.
// This provides an example of a system that creates a new entity.
func PlayerSpawnerSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.CreatePlayerMsg, msg.CreatePlayerResult](
		world,
		func(create message.TxData[msg.CreatePlayerMsg]) (msg.CreatePlayerResult, error) {
			id, err := createPlayer(world, create.Msg.Nickname)
			if err != nil {
				return msg.CreatePlayerResult{}, fmt.Errorf("error creating player: %w", err)
			}

			err = world.EmitEvent(map[string]any{
				"event": "new_player",
				"id":    id,
			})
			if err != nil {
				return msg.CreatePlayerResult{}, err
			}
			return msg.CreatePlayerResult{Success: true}, nil
		})
}

// createPlayer creates a player with the given name, and sets the player's HP to the InitialHP value.
// It also updates the PlayerNameToID global variable that maintains a mapping of player names to Entity IDs.
func createPlayer(world cardinal.WorldContext, name string) (types.EntityID, error) {
	id, err := cardinal.Create(world,
		comp.Player{Nickname: name},
		comp.Health{HP: InitialHP},
	)
	if err != nil {
		return 0, err
	}
	if PlayerNameToID == nil {
		PlayerNameToID = map[string]types.EntityID{}
	}
	PlayerNameToID[name] = id
	return id, nil
}
