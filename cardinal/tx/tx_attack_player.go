package tx

import (
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

type AttackPlayerMsg struct {
	TargetPlayerTag string `json:"target_player_tag"`
}

var AttackPlayer = ecs.NewTransactionType[AttackPlayerMsg]("attack-player")
