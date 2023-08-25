package main

import (
	"os"

	"github.com/argus-labs/starter-game-template/cardinal/utils"
	"pkg.world.dev/world-engine/cardinal/ecs"
)

type Config struct {
	CardinalPort string
	Mode         string
	RedisAddr    string
	RedisPass    string
}

func GetConfig() Config {
	mode := os.Getenv("REDIS_MODE")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	cardinalPort := os.Getenv("CARDINAL_PORT")

	return Config{
		CardinalPort: cardinalPort,
		Mode:         mode,
		RedisAddr:    redisAddr,
		RedisPass:    redisPassword,
	}
}

func NewWorld(cfg Config) *ecs.World {
	if cfg.Mode == "normal" {
		utils.NewWorld(cfg.RedisAddr, cfg.RedisPass)
	}
	return utils.NewEmbeddedWorld()
}
