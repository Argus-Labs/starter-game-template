package msg

import (
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

// TxHandler is a handler for `Tx` messages
// `Tx` is a type of message that performs a mutation on the state
type TxHandler struct {
	World *ecs.World
}
