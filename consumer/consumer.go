package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func handleErrorConsume(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	fmt.Println("Consume Application")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleErrorConsume(err, "Cannot Connect RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	handleErrorConsume(err, "Cannot Connect Chanel")
	defer ch.Close()
	msgs, err := ch.Consume("TestQueue", "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connect to our RabbitMQ Instance")
	fmt.Println(" [*] - waitingfor massages")
	<-forever
}
