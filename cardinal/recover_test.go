package main

import (
	"fmt"
	"testing"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/testutils"

	comp "github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/msg"
	"github.com/argus-labs/starter-game-template/cardinal/system"
)

// TestRecoveryOfNonECSState ensures data can be successfully saved and recovered across Cardinal restarts.
// When Cardinal restarts (and the ECS DB is still populated), in-memory go objects will be wiped out. These in-memory
// objects must be rebuilt using the data inside of ECS to ensure consistent behavior across cardinal restarts.
// NOTE: It's perfectly fine to use Cardinal's ECS storage to keep track of all your game state; this System and
// related test is included in the starter-game-template as an example for how to do in-memory object recovery and
// properly test it.
func TestRecoveryOfNonECSState(t *testing.T) {
	tf := testutils.NewTestFixture(t, nil)
	MustInitWorld(tf.World)

	tf.DoTick()

	// Make some players.
	for i := 0; i < 10; i++ {
		tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
			Nickname: fmt.Sprintf("player-%d", i),
		})

		tf.DoTick()
	}

	wCtx := cardinal.NewReadOnlyWorldContext(tf.World)
	// Make sure we can find those 10 players in our non-ecs data structure.
	for i := 0; i < 10; i++ {
		target := fmt.Sprintf("player-%d", i)
		id, ok := system.PlayerNameToID[target]
		if !ok {
			t.Fatalf("failed to find player %q in non-ecs storage", target)
		}
		p, err := cardinal.GetComponent[comp.Player](wCtx, id)
		if err != nil {
			t.Fatalf("failed to get player %q: %v", target, err)
		}
		if p.Nickname != target {
			t.Fatalf("player nickname does not match: got %q want %q", p.Nickname, target)
		}
	}

	// //////////////////////////////////////////////////////////////////////////////////////////////
	// Simulate Cardinal Restart: Save the now-populated redis DB for use in another text fixture. //
	// In addition, clear the in-memory go object (PlayerNameToID) to simulate a freshly restarted //
	// Cardinal process.                                                                           //
	// //////////////////////////////////////////////////////////////////////////////////////////////
	originalRedis := tf.Redis
	system.PlayerNameToID = nil

	// Create a new test fixture using the redis DB from the original test fixture.
	tf = testutils.NewTestFixture(t, originalRedis)
	MustInitWorld(tf.World)

	// The recover system must be given a chance to run.
	tf.DoTick()
	wCtx = cardinal.NewReadOnlyWorldContext(tf.World)
	// Make sure we can STILL find those 10 players in our non-ecs data structure.
	for i := 0; i < 10; i++ {
		target := fmt.Sprintf("player-%d", i)
		id, ok := system.PlayerNameToID[target]
		if !ok {
			t.Fatalf("failed to find player %q in non-ecs storage", target)
		}
		p, err := cardinal.GetComponent[comp.Player](wCtx, id)
		if err != nil {
			t.Fatalf("failed to get player %q: %v", target, err)
		}
		if p.Nickname != target {
			t.Fatalf("player lookup does not match expectation: got %q want %q", p.Nickname, target)
		}
	}
}
