package msg

import (
	msg2 "github.com/argus-labs/starter-game-template/msg/query"
	msg "github.com/argus-labs/starter-game-template/msg/tx"
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

// DEV NOTE: You don't have to edit this file.
// All tx and query definition should go to its own corresponding folder
// under msg/tx and msg/query respectively.
type MsgHandler struct {
	World *ecs.World
	Tx    msg.TxHandler
	Query msg2.QueryHandler
}

func NewMsgHandler(world *ecs.World) *MsgHandler {
	return &MsgHandler{
		World: world,
		Tx:    msg.TxHandler{World: world},
		Query: msg2.QueryHandler{World: world},
	}
}
