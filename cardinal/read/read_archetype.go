package read

import (
	"errors"
	"fmt"

	"github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/game"
	"pkg.world.dev/world-engine/cardinal"
)

type ArchetypeRequest struct {
	Label string `json:"label"`
}

type ArchetypeResponse struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

var Archetype = cardinal.NewQueryType[ArchetypeRequest, ArchetypeResponse]("archetype", queryArchetype)

func queryArchetype(wCtx cardinal.WorldContext, req ArchetypeRequest) (ArchetypeResponse, error) {
	var entities []interface{}
	var errs []error

	// Check if archetype label exist
	var archetype game.IArchetype
	var archetypeLabelIsFound = false
	for _, a := range game.Archetypes {
		if a.Label == req.Label {
			archetype = a
			archetypeLabelIsFound = true
			break
		}
	}
	if !archetypeLabelIsFound {
		return ArchetypeResponse{}, errors.New("invalid archetype label")
	}

	// Query for the archetype
	query, err := wCtx.NewSearch(cardinal.Exact(archetype.Components...))
	if err != nil {
		return ArchetypeResponse{}, err
	}
	err = query.Each(wCtx, func(id cardinal.EntityID) bool {
		entity := make(map[string]interface{})
		entity["id"] = id

		// Get all the component values
		healthComponent, err := cardinal.GetComponent[component.HealthComponent](wCtx, id)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to get component %s", component.HealthComponent{}.Name()))
			return true
		}
		entity[healthComponent.Name()] = healthComponent
		playerComponent, err := cardinal.GetComponent[component.PlayerComponent](wCtx, id)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to get component %s", component.HealthComponent{}.Name()))
			return true
		}
		entity[playerComponent.Name()] = playerComponent
		return true
	})
	if err != nil {
		return ArchetypeResponse{}, err
	}

	// Handle errors
	if len(errs) > 0 {
		return ArchetypeResponse{}, errors.Join(errs...)
	}

	return ArchetypeResponse{Label: req.Label, Value: entities}, nil
}
