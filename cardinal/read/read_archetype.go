package read

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/argus-labs/starter-game-template/cardinal/game"
	"reflect"

	"github.com/argus-labs/world-engine/cardinal/ecs"
	"github.com/argus-labs/world-engine/cardinal/ecs/filter"
	"github.com/argus-labs/world-engine/cardinal/ecs/storage"
)

type ArchetypeMsg struct {
	ArchetypeLabel string `json:"archetype_label"`
}

var Archetype = ecs.NewReadType[ArchetypeMsg]("archetype", queryArchetype)

func queryArchetype(world *ecs.World, m []byte) ([]byte, error) {
	var entities []interface{}
	var errs []error

	// Unmarshal msg json into struct
	msg := new(ArchetypeMsg)
	err := json.Unmarshal(m, msg)
	if err != nil {
		return nil, errors.New("failed to unmarshal archetype msg")
	}

	// Check if archetype label exist
	var archetype game.IArchetype
	var archetypeLabelIsFound = false
	for _, a := range game.Archetypes {
		if a.Label == msg.ArchetypeLabel {
			archetype = a
			archetypeLabelIsFound = true
			break
		}
	}
	if !archetypeLabelIsFound {
		return nil, errors.New("invalid archetype label")
	}

	// Query for the archetype
	query := ecs.NewQuery(filter.Exact(archetype.Components...))
	query.Each(world, func(id storage.EntityID) {
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
				return
			}

			entity[component.Name()] = value
		}

		entities = append(entities, entity)
	})
	// Handle the case where there is errors
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	// marshal entities to json string
	res, err := json.Marshal(entities)
	if err != nil {
		return nil, errors.New("failed to marshal archetype entities")
	}

	return res, nil
}
