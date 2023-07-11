package msg

import (
	"net/http"

	"github.com/argus-labs/starter-game-template/utils"
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

type CreatePlayerMsg struct {
	Tag string `json:"tag"`
}

var TxCreatePlayer = ecs.NewTransactionType[CreatePlayerMsg]()

// NOTE: We are going to be abstracting away this in the future, but for now
// you have to copy and paste this for each transaction type.
func (h *TxHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var msg CreatePlayerMsg
	err := utils.DecodeMsg[CreatePlayerMsg](r, &msg)
	if err != nil {
		utils.WriteError(w, "unable to decode create player tx", err)
		return
	}

	TxCreatePlayer.AddToQueue(h.World, msg)
	utils.WriteResult(w, "ok")
}
