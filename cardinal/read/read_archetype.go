package read

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/argus-labs/starter-game-template/cardinal/game"
	"pkg.world.dev/world-engine/cardinal/ecs"
	"pkg.world.dev/world-engine/cardinal/ecs/filter"
	"pkg.world.dev/world-engine/cardinal/ecs/storage"
)

type ArchetypeRequest struct {
	Label string `json:"label"`
}

type ArchetypeResponse struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

var Archetype = ecs.NewReadType[ArchetypeRequest, ArchetypeResponse]("archetype", queryArchetype)

func queryArchetype(world *ecs.World, req ArchetypeRequest) (ArchetypeResponse, error) {
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
	query := ecs.NewQuery(filter.Exact(archetype.Components...))
	query.Each(world, func(id storage.EntityID) bool {
		entity := make(map[string]interface{})
		entity["id"] = id

		// Get all the component values
		for _, component := range archetype.Components {
			// Call the component's Get method. This returns the component's value and an error
			in := []reflect.Value{reflect.ValueOf(world), reflect.ValueOf(id)}
			componentGetResult := reflect.ValueOf(component).MethodByName("Get").Call(in)

			value, err := componentGetResult[0].Interface(), componentGetResult[1].Interface()
			if err != nil {
				errs = append(errs, fmt.Errorf("failed to get component %s", component.Name()))
				return true
			}

			entity[component.Name()] = value
		}

		entities = append(entities, entity)
		return true
	})

	// Handle errors
	if len(errs) > 0 {
		return ArchetypeResponse{}, errors.Join(errs...)
	}

	return ArchetypeResponse{Label: req.Label, Value: entities}, nil
}
