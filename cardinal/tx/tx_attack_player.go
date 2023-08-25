package tx

import (
	"pkg.world.dev/world-engine/cardinal/ecs"
)

type AttackPlayerMsg struct {
	TargetNickname string `json:"target"`
}

type AttackPlayerMsgReply struct{}

var AttackPlayer = ecs.NewTransactionType[AttackPlayerMsg, AttackPlayerMsgReply]("attack-player")
