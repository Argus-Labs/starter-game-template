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
	DeployMode   string
}

func GetConfig() Config {
	mode := os.Getenv("REDIS_MODE")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	cardinalPort := os.Getenv("CARDINAL_PORT")
	deployMode := os.Getenv("DEPLOY_MODE")

	return Config{
		CardinalPort: cardinalPort,
		Mode:         mode,
		RedisAddr:    redisAddr,
		RedisPass:    redisPassword,
		DeployMode:   deployMode,
	}
}

func NewWorld(cfg Config) *ecs.World {
	if cfg.Mode == "normal" {
		return utils.NewWorld(cfg.RedisAddr, cfg.RedisPass, cfg.DeployMode)
	}
	return utils.NewEmbeddedWorld(cfg.DeployMode)
}
