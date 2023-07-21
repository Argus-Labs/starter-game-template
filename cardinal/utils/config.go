package utils

import (
	"github.com/argus-labs/world-engine/cardinal/ecs"
	"io"
	"net/http"
	"os"
)

type Config struct {
	CardinalPort string
	World        *ecs.World
}

var (
	EnvDFDeployMode  = os.Getenv("DF_DEPLOY_MODE")
	EnvRedisAddr     = os.Getenv("REDIS_ADDR")
	EnvRedisPassword = os.Getenv("REDIS_PASSWORD")
	EnvCardinalPort  = os.Getenv("CARDINAL_PORT")
)

func GetConfig() Config {
	if EnvDFDeployMode == "production" {
		return Config{
			CardinalPort: EnvCardinalPort,
			World:        NewProdWorld(EnvRedisAddr, EnvRedisPassword),
		}
	}
	return Config{
		CardinalPort: EnvCardinalPort,
		World:        NewDevWorld(EnvRedisAddr),
	}
}

func fetchGET(url string) ([]byte, error) {
	// Send GET request
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return body, err
}
