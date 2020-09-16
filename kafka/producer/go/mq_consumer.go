package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/segmentio/kafka-go"
	_ "github.com/segmentio/kafka-go/snappy"
	"log"
	"time"
)

func getKafkaReader(BrokersList []string, topic string, partition int) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   BrokersList,
		Topic:     topic,
		Partition: partition,
	})
}

func main() {
	flag.Parse()

	limit := 100000

	topic := "topic_bench_go"
	partition := 0

	kafkaURL := []string{"192.168.99.101:9092", "192.168.99.101:9093", "192.168.99.101:9094"}

	reader := getKafkaReader(kafkaURL, topic, partition)
	defer reader.Close()

	fmt.Println("start read ... !!")

	timestart := time.Now()

	for i := 0; i < limit; i++ {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("... message at topic:%v partition:%v offset:%v\n", m.Topic, m.Partition, m.Offset)
	}

	time.Sleep(time.Microsecond)

	timeEnd := time.Now()
	elapsed := timeEnd.Sub(timestart)
	log.Printf("Took %f seconds", elapsed.Seconds())

	throughut := float64(limit) / elapsed.Seconds()
	log.Printf("Throuphut %f/sec", throughut)
}
