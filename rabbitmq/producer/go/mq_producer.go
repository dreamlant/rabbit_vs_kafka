package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var username = flag.String("u", "guest", "RabbitMQ username")
	var password = flag.String("p", "guest", "RabbitMQ user password")
	var host = flag.String("h", "localhost", "RabbitMQ endpoint host or IP address")
	var throttling = flag.Int("t", 0, "Throttling per message (timeout between new message in seconds)")
	var limit = flag.Int("limit", 10, "Batch size (messages bunch count)")
	var messageSize = flag.Int("msize", 5, "Message size (bytes)")

	flag.Parse()

	body := RandStringBytes(*messageSize)

	timestart := time.Now()

	var connectEndpoint = fmt.Sprintf("amqp://%s:%s@%s:5672/", *username, *password, *host)
	log.Printf("Connecting to %s", connectEndpoint)
	conn, err := amqp.Dial(connectEndpoint)
	failOnError(err, "Failed to connect to RabbitMQ.")
	log.Printf("Connected.")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel.")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	var msg string
	for i := 0; i < *limit; i++ {
		msg = strconv.Itoa(i) + "___" + body
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(msg),
			})
		//log.Printf(" [x] Sent number %d --- %s", i, msg)
		failOnError(err, "Failed to publish a message")

		time.Sleep(time.Duration(*throttling) * time.Second)
	}

	timeEnd := time.Now()
	elapsed := timeEnd.Sub(timestart)
	log.Printf("Took %f seconds", elapsed.Seconds())

	throughut := float64(*limit) / elapsed.Seconds()
	log.Printf("Throughut %f/sec", throughut)
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = LetterBytes[rand.Intn(len(LetterBytes))]
	}
	return string(b)
}
