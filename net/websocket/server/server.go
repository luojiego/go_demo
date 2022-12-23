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

	for {
		trySend2Client(client)
		tryRecv(client)
	}
}

func trySend2Client(client *websocket.Conn) {
	var i int32 = 0
	for ; i < 10000; i++ {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.BigEndian, i); err != nil {
			fmt.Println(err)
		}
		// fmt.Printf("%+v\n", buf.Bytes())
		if err := client.WriteMessage(websocket.BinaryMessage, buf.Bytes()); err != nil {
			fmt.Printf("err: %s\n", err)
		}
	}
}

func tryRecv(client *websocket.Conn) error {
	var num int32 = 0
	_, buf, err := client.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return err
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
	return nil
}
