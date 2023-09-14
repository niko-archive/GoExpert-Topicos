package main

import (
	"fmt"
	"local/pkg/rabbitmq"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	QueueName = "main"
)

func main() {

	log.Default().Println("Consumer started")
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)

	go rabbitmq.Consume(QueueName, ch, out)

	for msg := range out {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}

}
