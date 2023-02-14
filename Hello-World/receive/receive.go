package main

import (
	"fmt"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

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
	if err != nil {
		log.Fatal(err)
	}

	msg, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := range msg {
			fmt.Printf("message received, '%s'\n", i.Body)
		}
	}()

	wg.Wait()
}
