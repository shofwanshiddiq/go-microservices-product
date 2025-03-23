package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQService struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewRabbitMQService(amqpURL string, queueName string) (*RabbitMQService, error) {
	// Connect to RabbitMQ
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		return nil, err
	}

	// Declare the queue
	q, err := ch.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err)
		return nil, err
	}

	return &RabbitMQService{
		connection: conn,
		channel:    ch,
		queue:      q,
	}, nil
}

func (r *RabbitMQService) ConsumeMessages(handler func([]byte) error) error {
	msgs, err := r.channel.Consume(
		r.queue.Name, // Queue name
		"",           // Consumer
		true,         // Auto-ack
		false,        // Exclusive
		false,        // No-local
		false,        // No-wait
		nil,          // Args
	)
	if err != nil {
		log.Printf("Failed to register a consumer: %v", err)
		return err
	}

	for msg := range msgs {
		if err := handler(msg.Body); err != nil {
			log.Printf("Failed to handle message: %v", err)
			// Optionally, you can add logic to requeue the message or handle the error
		}
	}

	return nil
}

func (r *RabbitMQService) Close() {
	r.channel.Close()
	r.connection.Close()
}
