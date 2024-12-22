package main

import (
	"os"

	"github.com/go-http-server/rabbitmq-socket/internal/rabbitmq"
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
	log.Info().Msg("loaded environment variables")

	rmqCli, err := rabbitmq.NewRabbitMQClient(envVars.RABBITMQ_URL)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to rabbitmq")
	}
	defer rmqCli.CloseChannel()
	log.Info().Msg("connected to rabbitmq")
}
