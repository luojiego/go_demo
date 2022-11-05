package main

import (
	"bytes"
	"encoding/binary"
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
	flag.Parse()
	u := url.URL{Scheme: *scheme, Host: *addr, Path: *path}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	// buf := make([]byte, 4)
	// for i := 0; i < len(buf); i++ {
	// 	buf[i] = byte(i)
	// }

	var i int32 = 0
	for ; i < 1000; i++ {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.BigEndian, i); err != nil {
			fmt.Println(err)
		}
		// fmt.Printf("%+v\n", buf.Bytes())
		if err := c.WriteMessage(websocket.BinaryMessage, buf.Bytes()); err != nil {
			fmt.Printf("err: %s\n", err)
		}
	}
}
