package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	fmt.Println("Golang RabbitMQ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Cannot Connect RabitMQ")
	defer conn.Close()

	fmt.Println("Successfully Connected	To Our RabbitMQ Instance")

	ch, err := conn.Channel()
	handleError(err, "Cannot Connect Chanel")
	defer ch.Close()

	q, err := ch.QueueDeclare("TestQueue", false, false, false, false, nil)
	handleError(err, "Cannot Connect Queue Declare")
	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)
	handleError(err, "Cannot Published Message Queue")
	fmt.Println("Successfully Published Message Queue")

}
