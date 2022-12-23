package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

var (
	topic     = "my-topic7"
	partition = 0
	address   = "192.168.196.37:9092"
)

//生产者
func producer() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(60 * time.Second))

	i := 0
	for {
		message := fmt.Sprintf("producer: %d", i+1)
		_, err := conn.WriteMessages(kafka.Message{Value: []byte(message)})
		if err != nil {
			panic(err)
		}

		fmt.Printf("write success: %s\n", message)
		time.Sleep(time.Second)
		i++
	}
}

//消费者
func consumer() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	batch := conn.ReadBatch(10, 1e6) //fetch 10KB min, 1MB max
	defer batch.Close()
	b := make([]byte, 10e3) //10kb max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		//fmt.Println("n: ", n)
		fmt.Println("consumer: ", string(b[:n]))
		time.Sleep(time.Second)
	}
}

func reader() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"192.168.196.37:9092"},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	r.SetOffset(42)
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}

func main() {
	//go producer()
	//go consumer()

	reader()

	//consumer()
	time.Sleep(60 * time.Second)
}
