package msg

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/argus-labs/starter-game-template/game"
	"github.com/argus-labs/starter-game-template/types"
	"github.com/argus-labs/starter-game-template/utils"
	"github.com/argus-labs/world-engine/cardinal/ecs"
	"github.com/argus-labs/world-engine/cardinal/ecs/filter"
	"github.com/argus-labs/world-engine/cardinal/ecs/storage"
)

type QueryArchetypeMsg struct {
	ArchetypeLabel string `json:"archetype_label"`
}

func (h *QueryHandler) Archetype(w http.ResponseWriter, r *http.Request) {
	var msg QueryArchetypeMsg
	err := utils.DecodeMsg[QueryArchetypeMsg](r, &msg)
	if err != nil {
		utils.WriteError(w, "unable to decode query archetype msg", err)
		return
	}

	entities, err := queryArchetype(h.World, msg)
	if err != nil {
		utils.WriteError(w, "failed to list archetype", err)
	} else {
		utils.WriteResult(w, entities)
	}
}

func queryArchetype(world *ecs.World, m QueryArchetypeMsg) ([]interface{}, error) {
	var entities []interface{}
	var errs []error

	var archetype types.IArchetype
	var archetypeLabelIsFound bool = false
	for _, a := range game.Archetypes {
		if a.Label == m.ArchetypeLabel {
			archetype = a
			archetypeLabelIsFound = true
			break
		}
	}
	if !archetypeLabelIsFound {
		return nil, errors.New("invalid archetype label")
	}

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

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return entities, nil
}
