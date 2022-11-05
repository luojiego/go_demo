package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

var (
	scheme = flag.String("scheme", "ws", "ws or wss")
	addr   = flag.String("addr", "127.0.0.1:8080", "http service address online test is hsahctest.letuinet.com")
	path   = flag.String("path", "/ws", "websocket path online test is /ws")
)

func main() {
	u := url.URL{Scheme: *scheme, Host: *addr, Path: *path}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	buf := make([]byte, 10)
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < 10; i++ {
		if err := c.WriteMessage(websocket.BinaryMessage, buf); err != nil {
			fmt.Printf("err: %s\n", err)
		}
	}
}
