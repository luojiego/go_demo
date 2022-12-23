package main

import (
	"fmt"
	"net"
)

const (
	addr = "127.0.0.1:8004"
)

func main() {
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		recvMessage(conn)
	}
}

func recvMessage(conn *net.UDPConn) {
	buf := make([]byte, 256)
	n, caddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Printf("read from udp err: %s", err)
		return
	}
	fmt.Printf("recv message: %s size: %d from: %s\n",
		string(buf), n, caddr.String())
}
