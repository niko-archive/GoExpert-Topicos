package main

import (
	"context"
	"fmt"
	"local/pkg/rabbitmq"
	"log"
)

const (
	QueueName = "main"
)

func main() {

	log.Default().Println("Producer started")
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	ctx := context.Background()

	for i := 0; i < 30000; i++ {

		msg := "Hello World: " + fmt.Sprintf("%d", i)
		rabbitmq.Publish(ctx, "amq.direct", "", msg, ch)
	}

}
