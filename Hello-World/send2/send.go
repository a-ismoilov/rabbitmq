package main

import (
	"context"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

const t = 5

func main() {
	conn, err := rabbitmq.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	q, err := ch.QueueDeclare(
		"queue",
		false,
		true,
		false,
		false,
		nil,
	)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*t)
	defer cancel()

	body := ", World"
	if err := ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		rabbitmq.Publishing{
			ContentType: "text/plain",
			Timestamp:   time.Now(),
			Body:        []byte(body),
		},
	); err != nil {
		log.Fatal(err)
	}

	log.Printf("message sent, '%s'\n", body)
}
