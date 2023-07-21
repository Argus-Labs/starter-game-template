package tx

import (
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

type CreatePlayerMsg struct {
	Tag string `json:"tag"`
}

var CreatePlayer = ecs.NewTransactionType[CreatePlayerMsg]("create-player")
