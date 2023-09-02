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

}
