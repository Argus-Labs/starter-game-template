package msg

import (
	"errors"
	"net/http"

	"github.com/argus-labs/starter-game-template/game"
	"github.com/argus-labs/starter-game-template/utils"
)

type QueryConstantMsg struct {
	ConstantLabel string `json:"constant_label"`
}

func (h *QueryHandler) Constant(w http.ResponseWriter, r *http.Request) {
	var msg QueryConstantMsg
	err := utils.DecodeMsg[QueryConstantMsg](r, &msg)
	if err != nil {
		utils.WriteError(w, "unable to decode query constant msg", err)
		return
	}

	entities, err := queryConstant(msg)
	if err != nil {
		utils.WriteError(w, "failed to list constant", err)
	} else {
		utils.WriteResult(w, entities)
	}
}

func queryConstant(m QueryConstantMsg) (interface{}, error) {
	// Handle all constants query
	if m.ConstantLabel == "all" {
		// Create a map of all constants and set it to be the value of result
		constants := make(map[string]interface{})
		for _, c := range game.ExposedConstants {
			constants[c.Label] = c.Value
		}
		return constants, nil
	}

	// Handle single constant query
	for _, constant := range game.ExposedConstants {
		if constant.Label == m.ConstantLabel {
			return constant.Value, nil
		}
	}

	return nil, errors.New("constant not found")
}
