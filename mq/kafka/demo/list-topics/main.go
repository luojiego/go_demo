package main

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

var (
	topic     = "t-2021-05-26"
	partition = 0
	address   = "192.168.196.19:9092"
)

func main() {
	conn, err := kafka.Dial("tcp", address)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
	fmt.Println("exited")
}
