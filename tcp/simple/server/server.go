package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("TCP", "localhost:9091")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		accept.Write([]byte("Hello"))
		accept.Close()
	}
}
