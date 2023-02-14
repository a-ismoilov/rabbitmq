package main

import (
	"context"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

const t = 5

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*t)
	defer cancel()
	conn, err := rabbitmq.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

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
	if err != nil {
		log.Fatal(err)
	}

	body := "Hello"
	if err := ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		rabbitmq.Publishing{
			ContentType: "text/plain",
			Timestamp:   time.Now(),
			Body:        []byte(body),
		}); err != nil {
		log.Fatal(err)
	}

	log.Printf("message sent, '%s'\n", body)
}
