package main

import (
	"bytes"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
	"time"
)

const t = 5

func main() {
	conn, err := rabbitmq.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()

	q, err := ch.QueueDeclare("hello",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := ch.Qos(1,
		0,
		false,
	); err != nil {
		log.Fatal(err)
	}

	msg, err := ch.Consume(q.Name,
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
			log.Printf("[x] Message Received := '%s'", i.Body)
			dotC := bytes.Count(i.Body, []byte("."))
			t := time.Duration(dotC)
			time.Sleep(t * time.Second)
			log.Println("Done")
			if err := i.Ack(false); err != nil {
				log.Fatal(err)
			}
		}
	}()
	log.Println("[X] Open to receive tasks [X]")

	wg.Wait()
}
