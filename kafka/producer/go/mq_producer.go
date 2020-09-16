package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"log"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func newKafkaWriter(BrokersList []string, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  BrokersList,
		Topic:    topic,
		Balancer: &kafka.RoundRobin{},
		//Balancer: &kafka.LeastBytes{},
		Async:            true,
		RequiredAcks:     1,
		CompressionCodec: snappy.NewCompressionCodec(),
	})
}

func main() {
	flag.Parse()

	limit := 100000

	// to produce messages
	topic := "topic_bench_go"
	//partition := 0

	kafkaURL := []string{"192.168.99.101:9092", "192.168.99.101:9093", "192.168.99.101:9094"}
	//msg := "test-async"
	msg := RandStringBytes(1024)

	writer := newKafkaWriter(kafkaURL, topic)

	defer writer.Close()

	fmt.Println("start producing ... !!")
	timestart := time.Now()

	for i := 0; i < limit; i++ {
		msg := kafka.Message{
			//Key:   []byte(fmt.Sprintf("Key-%d", i)),
			Value: []byte(msg),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		}
	}

	time.Sleep(time.Microsecond)

	timeEnd := time.Now()
	elapsed := timeEnd.Sub(timestart)
	log.Printf("Took %f seconds", elapsed.Seconds())

	throughut := float64(limit) / elapsed.Seconds()
	log.Printf("Throughut %f/sec", throughut)

	log.Print(writer.Stats())
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
