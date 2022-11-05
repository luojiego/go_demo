package main

import (
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
		_, bytes, err := client.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		fmt.Printf("len(bytes): %d\n", len(bytes))
	}
}
