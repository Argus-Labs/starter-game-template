package tx

import (
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

type AttackPlayerMsg struct {
	TargetNickname string `json:"target"`
}

var AttackPlayer = ecs.NewTransactionType[AttackPlayerMsg]("attack-player")
