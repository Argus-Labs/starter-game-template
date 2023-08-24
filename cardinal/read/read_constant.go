package read

import (
	"errors"

	"github.com/argus-labs/starter-game-template/cardinal/game"
	"pkg.world.dev/world-engine/cardinal/ecs"
)

type ConstantRequest struct {
	Label string `json:"label"`
}

type ConstantResponse struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

var Constant = ecs.NewReadType[ConstantRequest, ConstantResponse]("constant", queryConstant)

func queryConstant(_ *ecs.World, req ConstantRequest) (ConstantResponse, error) {
	var value interface{} = nil

	// Handle all constants query
	if req.Label == "all" {
		// Create a map of all constants
		constants := make(map[string]interface{})
		for _, c := range game.ExposedConstants {
			constants[c.Label] = c.Value
		}
		value = constants
	}

	// Handle single constant query
	for _, constant := range game.ExposedConstants {
		if constant.Label == req.Label {
			value = constant.Value
		}
	}

	if value == nil {
		return ConstantResponse{}, errors.New("constant not found")
	} else {
		return ConstantResponse{Label: req.Label, Value: value}, nil
	}
}
