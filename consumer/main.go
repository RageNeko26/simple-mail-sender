package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"fmt"
	"encoding/json"
)

type Payload struct {
	Email string `json:"email"`
	URL string `json:"url"`
}

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
			var newPayload Payload
			errUnmarshal := json.Unmarshal(d.Body, &newPayload)
			
			if errUnmarshal != nil {
				fmt.Println(errUnmarshal)
			}

			fmt.Println("Received message:", newPayload.Email)
			fmt.Println("Received message:", newPayload.URL)
		}
	}()

	fmt.Println("Successfully connect to RabbitMQ")
	fmt.Println("Waiting for message...")

	<-forever

	
}
