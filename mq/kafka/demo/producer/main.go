package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	topic     = "t-2021-05-26"
	partition = 0
	address   = "192.168.196.19:9092"
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
		message := "Roger message: " + strconv.FormatInt(int64(i+1), 10)
		_, err := conn.WriteMessages(kafka.Message{Value: []byte(message)})
		if err != nil {
			panic(err)
		}

		fmt.Printf("write success: %s\n", message)
		time.Sleep(time.Second)
		i++
	}
}

func main() {
	producer()
}
