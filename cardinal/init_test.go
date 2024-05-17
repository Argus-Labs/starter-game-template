package main

import (
	"fmt"
	"testing"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/testutils"
	"pkg.world.dev/world-engine/cardinal/types"

	"github.com/argus-labs/starter-game-template/cardinal/component"
)

// TestInitSystem_SpawnDefaultPlayersSystem_DefaultPlayersAreSpawned ensures a set of default players are created in the
// SpawnDefaultPlayersSystem. These players should only be created on tick 0.
func TestInitSystem_SpawnDefaultPlayersSystem_DefaultPlayersAreSpawned(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)

	tf.DoTick()
	// Do an extra tick to make sure the default players are only created once.
	tf.DoTick()

	wCtx := cardinal.NewReadOnlyWorldContext(tf.World)

	foundPlayers := map[string]bool{}
	searchErr := cardinal.NewSearch().Entity(filter.Contains(filter.Component[component.Health]())).
		Each(wCtx, func(id types.EntityID) bool {
			player, err := cardinal.GetComponent[component.Player](wCtx, id)
			if err != nil {
				t.Fatalf("failed to get player: %v", err)
			}
			health, err := cardinal.GetComponent[component.Health](wCtx, id)
			if err != nil {
				t.Fatalf("failed to get health: %v", err)
			}
			if health.HP < 100 {
				t.Fatalf("new player should have at least 100 health; got %v", health.HP)
			}
			foundPlayers[player.Nickname] = true
			return true
		})
	if searchErr != nil {
		t.Fatalf("failed to perform search: %v", searchErr)
	}
	if len(foundPlayers) != 10 {
		t.Fatalf("there should be 10 default players; got %v", foundPlayers)
	}
	for i := 0; i < 10; i++ {
		wantName := fmt.Sprintf("default-%d", i)
		if !foundPlayers[wantName] {
			t.Fatalf("could not find player %q in %v", wantName, foundPlayers)
		}
	}
}
