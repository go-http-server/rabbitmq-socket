package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQClient(rabbitmqURL string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQClient{
		conn:    conn,
		channel: channel,
	}, nil
}

func (rc *RabbitMQClient) CloseChannel() {
	log.Info().Msg("Closing channel...")
	log.Info().Msg("Closing connection...")
	defer rc.channel.Close()
	defer rc.conn.Close()
}
