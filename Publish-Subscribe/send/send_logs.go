package main

import (
	"context"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

const (
	t = 5
)

func main() {
	conn, err := rabbitmq.Dial("amqp://guest:guest@localhost:5672")
	errFatal(err)

	ch, err := conn.Channel()
	errFatal(err)
	defer ch.Close()

	err = ch.ExchangeDeclare("logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	errFatal(err)
	ctx, cancel := context.WithTimeout(context.TODO(), t*time.Second)
	defer cancel()

	b := body(os.Args)

	err = ch.PublishWithContext(ctx,
		"logs",
		"",
		false,
		false,
		rabbitmq.Publishing{
			ContentType: "text/plain",
			Timestamp:   time.Now(),
			Body:        []byte(b),
		},
	)
	errFatal(err)
	log.Println("[x] Message sent")
}

func errFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func body(args []string) string {
	result := ""
	if len(args) < 2 {
		log.Fatal("Should be in this form: go run send_logs.go ...")
	} else {
		for i, arg := range args {
			if i == 0 {
				continue
			}
			if i == len(args)-1 {
				result += arg
			} else {
				result += arg + " "
			}
		}
	}

	return result
}
