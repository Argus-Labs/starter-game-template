package tx

import (
	"pkg.world.dev/world-engine/cardinal"
)

type AttackPlayerMsg struct {
	TargetNickname string `json:"target"`
}

type AttackPlayerMsgReply struct{}

var AttackPlayer = cardinal.NewTransactionType[AttackPlayerMsg, AttackPlayerMsgReply]("attack-player")
