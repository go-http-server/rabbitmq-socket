package rabbitmq

import (
	"context"

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

// WARN: durable is parameter decide queue won't be delete after rabbitmq restart. autoDelete -> delete queue when no CONSUMER listen that.
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

// CreateBinding will bind the current channel to the given exchange using the using routing key provided
func (rc *RabbitMQClient) QueueBind(queueName, bindingName, exchangeName string) error {
	// leaving noWait false, having noWait set to false will make the channel return an error if its fails to binding
	return rc.channel.QueueBind(queueName, bindingName, exchangeName, false, nil)
}

func (rc *RabbitMQClient) PublishEvent(ctx context.Context, exchangeName, routingKey string, options amqp.Publishing) error {
	confirm, err := rc.channel.PublishWithDeferredConfirmWithContext(ctx, exchangeName, routingKey, true, false, options)
	if err != nil {
		return err
	}
	isConfirmed, err := confirm.WaitContext(ctx)
	if err != nil {
		return err
	}
	log.Printf("publish event with deferred confirm with context: %t", isConfirmed)
	return nil
}

func (rc *RabbitMQClient) CloseChannel() {
	log.Info().Msg("closing channel...")
	err := rc.channel.Close()
	if err != nil {
		log.Error().Err(err).Msg("cannot close channel")
	}
	log.Info().Msg("closed channel")
}
