package tx

import (
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

type AttackPlayerMsg struct {
	TargetNickname string `json:"target"`
}

type AttackPlayerMsgReply struct{}

var AttackPlayer = ecs.NewTransactionType[AttackPlayerMsg, AttackPlayerMsgReply]("attack-player")
