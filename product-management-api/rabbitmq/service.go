// rabbitmq/service.go
package rabbitmq

import (
	"encoding/json"
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

func (r *RabbitMQService) PublishMessage(message interface{}) error {
	// Serialize the message to JSON
	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to serialize message: %v", err)
		return err
	}

	// Publish the message to the queue
	err = r.channel.Publish(
		"",           // Exchange
		r.queue.Name, // Routing key (queue name)
		false,        // Mandatory
		false,        // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageJSON,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	log.Println("Message published to RabbitMQ:", message)
	return nil
}

func (r *RabbitMQService) Close() {
	r.channel.Close()
	r.connection.Close()
}
