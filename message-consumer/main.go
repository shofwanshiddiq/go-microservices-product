package main

import (
	"log"
	"message-consumer/rabbitmq"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	amqpURL := "amqp://user:password@rabbitmq:5672/"
	queueName := "your_queue_name"

	rabbitMQService, err := rabbitmq.NewRabbitMQService(amqpURL, queueName)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	defer rabbitMQService.Close()

	// Handle incoming messages
	go func() {
		if err := rabbitMQService.ConsumeMessages(handleMessage); err != nil {
			log.Fatalf("Failed to consume messages: %v", err)
		}
	}()

	// Wait for termination signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down message consumer...")
}

func handleMessage(message []byte) error {
	// Handle the incoming message (e.g., log it)
	log.Printf("Received message: %s", message)
	return nil
}
