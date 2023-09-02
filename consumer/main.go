package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"fmt"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	
	if err != nil {
		fmt.Println("Failed to connect instance RabbitMQ")
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println("Failed to open channel!")
	}

	defer ch.Close()

	msg, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msg {
			fmt.Println("Received message:", string(d.Body))
		}
	}()

	fmt.Println("Successfully connect to RabbitMQ")
	fmt.Println("Waiting for message...")

	<-forever

	
}
