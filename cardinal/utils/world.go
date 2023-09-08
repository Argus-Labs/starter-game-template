package utils

import (
	"time"

	"github.com/rs/zerolog/log"

	"pkg.world.dev/world-engine/cardinal"
)

const loopInterval = time.Second

func getOptions(deployMode string) []cardinal.WorldOption {
	var options []cardinal.WorldOption
	if deployMode == "development" {
		options = append(options, cardinal.WithPrettyLog())
	}
	options = append(options, cardinal.WithLoopInterval(loopInterval))

	// TODO: When launching to production, you should enable signature verification.
	options = append(options, cardinal.WithDisableSignatureVerification())
	return options
}

// NewWorld is the recommended way to run the game
func NewWorld(addr string, password string, deployMode string) *cardinal.World {
	log.Log().Msg("Running in normal mode, using external Redis")
	if addr == "" {
		log.Log().Msg("Redis address is not set, using fallback - localhost:6379")
		addr = "localhost:6379"
	}
	if password == "" {
		log.Log().Msg("Redis password is not set, make sure to set up redis with password in prod")
		password = ""
	}
	world, err := cardinal.NewWorld(addr, password, getOptions(deployMode)...)
	if err != nil {
		panic(err)
	}

	return world
}

// NewEmbeddedWorld is the most convenient way to run the game
// because it doesn't require spinning up Redis in a container.
// It runs a Redis server as a part of the Go process.
// NOTE: worlds with embedded redis are incompatible with Cardinal Editor.
func NewEmbeddedWorld(deployMode string) *cardinal.World {
	log.Log().Msg("Running in embedded mode, using embedded miniredis")
	world, err := cardinal.NewMockWorld(getOptions(deployMode)...)
	if err != nil {
		log.Fatal().Err(err)
	}
	return world
}
