package system

import (
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "github.com/argus-labs/starter-game-template/cardinal/component"
)

// PlayerNameToID maps a player's name to the relevant entity ID. This information is stored in a global variable which
// is NOT a part of the ECS data model. Care must be taken to 1) ensure the data remains consistent with the data
// contained in the ECS data model and 2) the data is recovered from ECS when Cardinal restarts.
var PlayerNameToID map[string]types.EntityID

// RecoverPlayersSystem demonstrates a recovery system pattern. The system will be executed every game tick, however
// the main logic of the system will only be executed once each time Cardinal restarts. This pattern can be used to
// maintain non-ECS state. This recovery step is required to ensure the in-memory data (which is wiped out when Cardinal
// restarts) is maintained across cardinal restarts. These recovery systems should be registered BEFORE any other
// systems that may depend on the in-memory information.
func RecoverPlayersSystem(world cardinal.WorldContext) error {
	if PlayerNameToID != nil {
		return nil
	}
	// Ensure this recovery step only happens once each time Cardinal is started.
	PlayerNameToID = map[string]types.EntityID{}

	return cardinal.NewSearch(world, filter.Contains(comp.Player{})).Each(func(id types.EntityID) bool {
		player, err := cardinal.GetComponent[comp.Player](world, id)
		if err != nil {
			panic(err)
		}
		PlayerNameToID[player.Nickname] = id
		return true
	})
}
