package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	example "github.com/micro/examples/server/proto/example"
	"io"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:9091")
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
		fmt.Println("connect ", accept.RemoteAddr())

		go processor(accept)
	}
}

func processor(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		read, err := conn.Read(buf)
		if err != nil {
			log.Printf("read err: %s\n", err)
			if err == io.EOF {
				break
			}
			continue
		}

		request := &example.Request{}
		if err := proto.Unmarshal(buf[:read], request); err != nil {

			log.Printf("proto unmarshal err: %s\n", err)
			continue
		}
		//accept.Write([]byte("Hello"))
		fmt.Println(request)
		response := &example.Response{
			Msg: "Hello" + request.Name,
		}

		res, _ := proto.Marshal(response)
		conn.Write(res)
	}
}
