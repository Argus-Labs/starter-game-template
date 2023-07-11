package utils

import "github.com/rs/zerolog/log"

func Must(err error) {
	if err != nil {
		log.Fatal().Err(err)
	}
}
