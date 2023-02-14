package main

import (
	"fmt"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
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

	q, err := ch.QueueDeclare(
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	errFatal(err)
	err = ch.QueueBind(q.Name,
		"",
		"logs",
		false,
		nil,
	)
	errFatal(err)

	msg, err := ch.Consume(q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	errFatal(err)
	log.Println("[x] Waiting for the messages ...")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range msg {
			fmt.Println(string(i.Body))
		}
	}()

	wg.Wait()
}

func errFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
