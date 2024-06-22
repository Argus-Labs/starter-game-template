package main

import (
	"testing"

	"gotest.tools/v3/assert"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/receipt"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	"github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/msg"
)

const (
	attackMsgName = "game.attack-player"
	createMsgName = "game.create-player"
)

// TestSystem_AttackSystem_ErrorWhenTargetDoesNotExist ensures the attack message results in an error when the given
// target does not exist. Note, message errors are stored in receipts; they are NOT returned from the relevant system.
func TestSystem_AttackSystem_ErrorWhenTargetDoesNotExist(t *testing.T) {
	tf := cardinal.NewTestFixture(t, nil)
	MustInitWorld(tf.World)

	txHash := tf.AddTransaction(getAttackMsgID(t, tf.World), msg.AttackPlayerMsg{
		TargetNickname: "does-not-exist",
	})

	tf.DoTick()

	gotReceipt := getReceiptFromPastTick(t, tf.World, txHash)
	if len(gotReceipt.Errs) == 0 {
		t.Fatal("expected error when target does not exist")
	}
}

// TestSystem_PlayerSpawnerSystem_CanCreatePlayer ensures the CreatePlayer message can be used to create a new player
// with the default amount of health. cardinal.NewSearch is used to find the newly created player.
func TestSystem_PlayerSpawnerSystem_CanCreatePlayer(t *testing.T) {
	tf := cardinal.NewTestFixture(t, nil)
	MustInitWorld(tf.World)

	const nickname = "jeff"
	createTxHash := tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		Nickname: nickname,
	})
	tf.DoTick()

	// Make sure the player creation was successful
	createReceipt := getReceiptFromPastTick(t, tf.World, createTxHash)
	if errs := createReceipt.Errs; len(errs) > 0 {
		t.Fatalf("expected 0 errors when creating a player, got %v", errs)
	}

	// Make sure the newly created player has 100 health
	wCtx := cardinal.NewReadOnlyWorldContext(tf.World)
	// This search demonstrates the use of a "Where" clause, which limits the search results to only the entity IDs
	// that end up returning true from the anonymous function. In this case, we're looking for a specific nickname.
	acc := make([]types.EntityID, 0)
	err := cardinal.NewSearch().Entity(filter.All()).Each(wCtx, func(id types.EntityID) bool {
		player, err := cardinal.GetComponent[component.Player](wCtx, id)
		if err != nil {
			t.Fatalf("failed to get player component: %v", err)
		}
		if player.Nickname == nickname {
			acc = append(acc, id)
			return false
		}
		return true
	})
	assert.NilError(t, err)
	assert.Equal(t, len(acc), 1)
	id := acc[0]

	health, err := cardinal.GetComponent[component.Health](wCtx, id)
	if err != nil {
		t.Fatalf("failed to find entity ID: %v", err)
	}
	if health.HP != 100 {
		t.Fatalf("a newly created player should have 100 health; got %v", health.HP)
	}
}

// TestSystem_AttackSystem_AttackingTargetReducesTheirHealth ensures an attack message can find an existing target the
// reduce the target's health.
func TestSystem_AttackSystem_AttackingTargetReducesTheirHealth(t *testing.T) {
	tf := cardinal.NewTestFixture(t, nil)
	MustInitWorld(tf.World)

	const target = "jeff"

	// Create an initial player
	_ = tf.AddTransaction(getCreateMsgID(t, tf.World), msg.CreatePlayerMsg{
		Nickname: target,
	})
	tf.DoTick()

	// Attack the player
	attackTxHash := tf.AddTransaction(getAttackMsgID(t, tf.World), msg.AttackPlayerMsg{
		TargetNickname: target,
	})
	tf.DoTick()

	// Make sure attack was successful
	attackReceipt := getReceiptFromPastTick(t, tf.World, attackTxHash)
	if errs := attackReceipt.Errs; len(errs) > 0 {
		t.Fatalf("expected no errors when attacking a player; got %v", errs)
	}

	// Find the attacked player and check their health.
	wCtx := cardinal.NewReadOnlyWorldContext(tf.World)
	var found bool
	// This search demonstrates the "Each" pattern. Every entity ID is considered, and as long as the anonymous
	// function return true, the search will continue.
	searchErr := cardinal.NewSearch().Entity(filter.All()).Each(wCtx, func(id types.EntityID) bool {
		player, err := cardinal.GetComponent[component.Player](wCtx, id)
		if err != nil {
			t.Fatalf("failed to get player component for %v", id)
		}
		if player.Nickname != target {
			return true
		}
		// The player's nickname matches the target. This is the player we care about.
		found = true
		health, err := cardinal.GetComponent[component.Health](wCtx, id)
		if err != nil {
			t.Fatalf("failed to get health component for %v", id)
		}
		// The target started with 100 HP, -10 for the attack, +1 for regen
		if health.HP != 91 {
			t.Fatalf("attack target should end up with 91 hp, got %v", health.HP)
		}

		return false
	})
	if searchErr != nil {
		t.Fatalf("error when performing search: %v", searchErr)
	}
	if !found {
		t.Fatalf("failed to find target %q", target)
	}
}

func getCreateMsgID(t *testing.T, world *cardinal.World) types.MessageID {
	return getMsgID(t, world, createMsgName)
}

func getAttackMsgID(t *testing.T, world *cardinal.World) types.MessageID {
	return getMsgID(t, world, attackMsgName)
}

func getMsgID(t *testing.T, world *cardinal.World, fullName string) types.MessageID {
	msg, ok := world.GetMessageByFullName(fullName)
	if !ok {
		t.Fatalf("failed to get %q message", fullName)
	}
	return msg.ID()
}

// getReceiptFromPastTick search past ticks for a txHash that matches the given txHash. An error will be returned if
// the txHash cannot be found in Cardinal's history.
func getReceiptFromPastTick(t *testing.T, world *cardinal.World, txHash types.TxHash) receipt.Receipt {
	tick := world.CurrentTick()
	for {
		tick--
		receipts, err := world.GetTransactionReceiptsForTick(tick)
		if err != nil {
			t.Fatal(err)
		}
		for _, r := range receipts {
			if r.TxHash == txHash {
				return r
			}
		}
	}
}
