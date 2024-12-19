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

func (rc *RabbitMQClient) CreateQueue(queueName string, durable, autoDelete bool, argsTable amqp.Table) (amqp.Queue, error) {
	q, err := rc.channel.QueueDeclare(
		queueName,
		durable,    // if true then queue will be hold after rabbitmq restart.
		autoDelete, // if fasle then queue will be hold although queue not have CONSUMER listen that.
		false,      // exclusive is true then allow only current connection connect to queue.
		false,      // noWait true when system no needed response immediate
		argsTable,
	)
	if err != nil {
		return amqp.Queue{}, err
	}

	return q, nil
}

func (rc *RabbitMQClient) CloseChannel() {
	log.Info().Msg("Closing channel...")
	defer rc.channel.Close()
}
