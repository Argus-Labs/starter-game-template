package tx

import (
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

type CreatePlayerMsg struct {
	Nickname string `json:"nickname"`
}

var CreatePlayer = ecs.NewTransactionType[CreatePlayerMsg]("create-player")
