package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

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

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("Hello from RabbitMQ"),
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Publishing Message to Queue is success!")
}
