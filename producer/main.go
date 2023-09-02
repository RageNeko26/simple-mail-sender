package main

import (
	"fmt"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Payload struct {
	Email string `json:"email"`
	URL string `json:"url"`
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	defer conn.Close()

	fmt.Println("Connection to RabbitMQ is success!")

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)

	if err != nil {
		panic(err)
	}

	newPayload := Payload{
		Email: "test@tester.com",
		URL: "http://contoh.com",
	}

	marshalled, _ := json.Marshal(newPayload) 

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: marshalled,
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Publishing Message to Queue is success!")
}
