package main

import (
	"github.com/argus-labs/starter-game-template/cardinal/utils"
	"github.com/argus-labs/world-engine/cardinal/ecs"
	"os"
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
