package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	topic     = "t-2021-05-26"
	partition = 0
	address   = "192.168.196.19:9092"
)

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
			fmt.Printf("err: %s\n", err)
			break
		}
		//fmt.Println("n: ", n)
		fmt.Println("consumer: ", string(b[:n]))
		time.Sleep(time.Microsecond * 100)
	}
}

func reader() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{address},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	r.SetOffset(42)
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("ReadMessage err: %s\n", err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}

func main() {
	// 消费数据
	// consumer()
	reader()
	fmt.Println("exited")
}
