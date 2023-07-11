package msg

import (
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

// QueryHandler is a handler for `Query` messages
// `Query` is a type of message that does not perform mutation on the state
type QueryHandler struct {
	World *ecs.World
}
