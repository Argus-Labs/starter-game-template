package tx

import (
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

type CreatePlayerMsg struct {
	Nickname string `json:"nickname"`
}

type CreatePlayerMsgReply struct {
	Success bool `json:"success"`
}

var CreatePlayer = ecs.NewTransactionType[CreatePlayerMsg, CreatePlayerMsgReply]("create-player")
