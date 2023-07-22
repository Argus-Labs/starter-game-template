package read

import (
	"encoding/json"
	"errors"
	"github.com/argus-labs/starter-game-template/cardinal/game"
	"github.com/argus-labs/world-engine/cardinal/ecs"
)

type ConstantMsg struct {
	ConstantLabel string `json:"constant_label"`
}

var Constant = ecs.NewReadType[ConstantMsg]("constant", queryConstant)

func queryConstant(_ *ecs.World, m []byte) ([]byte, error) {
	// Unmarshal json into a ReadConstantMsg
	var msg ConstantMsg
	err := json.Unmarshal(m, &msg)
	if err != nil {
		return nil, err
	}

	// Handle all constants query
	if msg.ConstantLabel == "all" {
		// Create a map of all constants and set it to be the value of result
		constants := make(map[string]interface{})
		for _, c := range game.ExposedConstants {
			constants[c.Label] = c.Value
		}

		// Marshal the map into json
		res, err := json.Marshal(constants)
		if err != nil {
			return nil, err
		}

		return res, nil
	}

	// Handle single constant query
	for _, constant := range game.ExposedConstants {
		if constant.Label == msg.ConstantLabel {
			// Marshal the constant into json
			res, err := json.Marshal(constant.Value)
			if err != nil {
				return nil, err
			}

			return res, nil
		}
	}

	return nil, errors.New("constant not found")
}
