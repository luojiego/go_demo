package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, err2 := net.InterfaceAddrs()
	if err2 != nil {
		panic(err2)
	}
	for _, addr := range addrs {
		fmt.Println(addr.String())
	}
}
