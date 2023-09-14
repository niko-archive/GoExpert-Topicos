package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	conx, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	ch, err := conx.Channel()
	if err != nil {
		panic(err)
	}

	return ch, nil
}

func Consume(queue string, ch *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := ch.Consume(
		queue,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

func Publish(ctx context.Context, exchange string, name string, msg string, ch *amqp.Channel) error {
	err := ch.PublishWithContext(
		ctx,
		exchange,
		name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
