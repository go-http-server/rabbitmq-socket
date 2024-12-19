package main

import (
	"os"

	"github.com/go-http-server/rabbitmq-socket/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	envVars, err := utils.LoadEnvironmentVariables("./")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load environment variables")
	}
	if envVars.MODE_ENV == utils.DEVELOPMENT_MODE {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	log.Info().Str("HTTP_SERVER_ADDRESS", envVars.HTTP_SERVER_ADDRESS).Msg("loaded environment variables")
}
