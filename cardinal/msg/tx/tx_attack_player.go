package msg

import (
	"net/http"

	"github.com/argus-labs/starter-game-template/utils"
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

type AttackPlayerMsg struct {
	TargetPlayerTag string `json:"target_player_tag"`
}

var TxAttackPlayer = ecs.NewTransactionType[AttackPlayerMsg]()

// NOTE: We are going to be abstracting away this in the future, but for now
// you have to copy and paste this for each transaction type.
func (h *TxHandler) AttackPlayer(w http.ResponseWriter, r *http.Request) {
	var msg AttackPlayerMsg
	err := utils.DecodeMsg[AttackPlayerMsg](r, &msg)
	if err != nil {
		utils.WriteError(w, "unable to decode attack player tx", err)
		return
	}
	TxAttackPlayer.AddToQueue(h.World, msg)
	utils.WriteResult(w, "ok")
}
