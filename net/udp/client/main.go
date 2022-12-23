package main

import (
	"encoding/json"
	"fmt"
	"net"
)

const (
	addr = "127.0.0.1:8004"
)

type Data struct {
	UserId int    `json:"user_id,omitempty"`
	Source string `json:"source,omitempty"`
}

func main() {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	d := &Data{
		UserId: 10024,
		Source: "test",
	}
	result, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	n, err := conn.Write(result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("write size: %d\n", n)
}
