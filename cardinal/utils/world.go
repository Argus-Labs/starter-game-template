package utils

import (
	"github.com/rs/zerolog/log"
	"pkg.world.dev/world-engine/cardinal"
)

// NewWorld is the recommended way to run the game
func NewWorld(addr string, password string, deployMode string, options ...cardinal.WorldOption) *cardinal.World {
	log.Log().Msg("Running in normal mode, using external Redis")
	if addr == "" {
		log.Log().Msg("Redis address is not set, using fallback - localhost:6379")
		addr = "localhost:6379"
	}
	if password == "" {
		log.Log().Msg("Redis password is not set, make sure to set up redis with password in prod")
		password = ""
	}
	if deployMode == "development" {
		options = append(options, cardinal.WithPrettyLog())
	}

	res, err := cardinal.NewWorld(addr, password, options...)
	if err != nil {
		panic(err)
	}
	return res

	//return world
}

// NewEmbeddedWorld is the most convenient way to run the game
// because it doesn't require spinning up Redis in a container.
// It runs a Redis server as a part of the Go process.
// NOTE: worlds with embedded redis are incompatible with Cardinal Editor.
func NewEmbeddedWorld(deployMode string) *cardinal.World {
	log.Log().Msg("Running in embedded mode, using embedded miniredis")
	options := make([]cardinal.WorldOption, 0)
	if deployMode == "development" {
		options = append(options, cardinal.WithPrettyLog())
	}
	res, err := cardinal.NewMockWorld(options...)
	if err != nil {
		panic(err)
	}
	return res
}
