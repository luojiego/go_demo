package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

func main() {
	r := gin.Default()
	r.GET("/ws", root)
	r.GET("/test", test)
	r.Run(":8080")
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": true})
}

func root(c *gin.Context) {
	client, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer client.Close()

	var num int32 = 0
	for {
		_, buf, err := client.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		// fmt.Printf("buf: %+v\n", buf)
		reader := bytes.NewReader(buf)
		var val int32
		binary.Read(reader, binary.BigEndian, &val)
		if num != val {
			fmt.Printf("req: %d but send: %d\n", num, val)
		} else {
			num++
		}
		// fmt.Printf("recv val: %d\n", val)
	}
}
