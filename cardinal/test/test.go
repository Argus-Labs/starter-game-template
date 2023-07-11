package test

import (
	"github.com/argus-labs/starter-game-template/component"
	tx "github.com/argus-labs/starter-game-template/msg/tx"
	"github.com/argus-labs/starter-game-template/system"
	"github.com/argus-labs/starter-game-template/utils"
	"github.com/argus-labs/world-engine/cardinal/ecs"
	"github.com/argus-labs/world-engine/cardinal/ecs/storage"
)

// Miscellaneous test utilities
func ScaffoldTestWorld() *ecs.World {
	world := utils.NewInmemWorld()

	utils.Must(world.RegisterTransactions(
		tx.TxCreatePlayer,
	))

	world.AddSystem(system.PlayerSpawnerSystem)

	utils.Must(world.LoadGameState())

	return world
}

func CreatePlayer(world *ecs.World, tag string) (storage.EntityID, component.PlayerComponent) {
	entityId, _ := world.Create(component.Player)
	playerComp := component.PlayerComponent{
		Tag: tag,
	}
	component.Player.Set(world, entityId, playerComp)
	return entityId, playerComp
}
