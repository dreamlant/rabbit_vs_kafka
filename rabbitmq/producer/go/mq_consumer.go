package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var username = flag.String("u", "guest", "RabbitMQ username")
	var password = flag.String("p", "guest", "RabbitMQ user password")
	var host = flag.String("h", "localhost", "RabbitMQ endpoint host or IP address")
	var limit = flag.Int("limit", 10, "Batch size (messages bunch count)")

	flag.Parse()

	var connectEndpoint = fmt.Sprintf("amqp://%s:%s@%s:5672/", *username, *password, *host)
	log.Printf("Connecting to %s", connectEndpoint)

	conn, err := amqp.Dial(connectEndpoint)
	failOnError(err, "Failed to connect to RabbitMQ.")
	log.Printf("Connected.")
	defer conn.Close()

	timestart := time.Now()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel.")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue.")

	for i := 0; i < 5000; i++ {
		_, err = ch.Consume(
			queue.Name, // name
			"",         // consumerTag,
			true,       // ack auto
			false,      // exclusive
			false,      // noLocal
			false,      // noWait
			nil,        // arguments
		)
	}

	timeEnd := time.Now()
	elapsed := timeEnd.Sub(timestart)
	log.Printf("Took %f seconds", elapsed.Seconds())

	throughut := float64(*limit) / elapsed.Seconds()
	log.Printf("Throughut %f/sec", throughut)
}
