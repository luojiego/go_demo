package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	example "github.com/micro/examples/server/proto/example"
	"log"
	"net"
	"time"
)

func main() {
	begin := time.Now()
	conn, err := net.Dial("tcp", "localhost:9091")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for i := 0; i < 100; i++ {
		request := example.Request{
			Name: "Roger",
		}
		result, err := proto.Marshal(&request)
		if err != nil {
			continue
		}
		_, err = conn.Write(result)
		if err != nil {
			panic(err)
		}
		//log.Println("write len: ", write)

		buf := make([]byte, 1024)
		len, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}

		response := &example.Response{}
		proto.Unmarshal(buf[:len], response)
		log.Println(response)
	}
	fmt.Printf("use %0.2fs\n", time.Since(begin).Seconds())
}
