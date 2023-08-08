package main

import (
	"github.com/argus-labs/starter-game-template/cardinal/utils"
	"github.com/argus-labs/world-engine/cardinal/ecs"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	CardinalPort string
	World        *ecs.World
}

var (
	EnvRedisMode     = os.Getenv("REDIS_MODE")
	EnvRedisAddr     = os.Getenv("REDIS_ADDR")
	EnvRedisPassword = os.Getenv("REDIS_PASSWORD")
	EnvCardinalPort  = os.Getenv("CARDINAL_PORT")
)

func GetConfig() Config {
	if EnvRedisMode == "normal" {
		return Config{
			CardinalPort: EnvCardinalPort,
			World:        utils.NewWorld(EnvRedisAddr, EnvRedisPassword),
		}
	} else if EnvRedisMode == "embedded" {
		return Config{
			CardinalPort: EnvCardinalPort,
			World:        utils.NewEmbeddedWorld(),
		}
	} else {
		log.Log().Msg("REDIS_MODE is not set, using fallback - embedded")
		return Config{
			CardinalPort: EnvCardinalPort,
			World:        utils.NewEmbeddedWorld(),
		}
	}
}
