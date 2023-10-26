package tx

import (
	"pkg.world.dev/world-engine/cardinal"
)

type CreatePlayerMsg struct {
	Nickname string `json:"nickname"`
}

type CreatePlayerMsgReply struct {
	Success bool `json:"success"`
}

var CreatePlayer = cardinal.NewTransactionType[CreatePlayerMsg, CreatePlayerMsgReply]("create-player")
