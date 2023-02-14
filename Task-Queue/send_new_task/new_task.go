package main

import (
	"context"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strings"
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
		"durable",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	body := bodyFrom(os.Args)

	if err := ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		rabbitmq.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: rabbitmq.Persistent,
			Timestamp:    time.Now(),
			Body:         []byte(body),
		}); err != nil {
		log.Fatal(err)
	}

	log.Printf("[x] Message Sent := '%s'", body)
}

func bodyFrom(args []string) string {
	if (len(args) < 2) || os.Args[1] == "" {
		return "hello"
	}

	return strings.Join(args[1:], " ")
}
