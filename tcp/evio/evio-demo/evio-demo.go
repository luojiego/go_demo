package main

import (
	"bytes"
	"fmt"
	"github.com/tidwall/evio"
	"log"
	"time"
)

type conn struct {
	is   evio.InputStream
	addr string
}

var (
	s       []evio.Conn
	srv     evio.Server
	message string
)

func main() {
	events := evio.Events{
		NumLoops: -1,
	}

	events.Closed = func(c evio.Conn, err error) (action evio.Action) {
		fmt.Printf("disconnect connect: %s\n", c.RemoteAddr())
		message = fmt.Sprintf("%s disconnected\n", c.RemoteAddr())
		for _, v := range s {
			v.Wake()
		}
		return
	}
	/*events.Tick = func() (delay time.Duration, action evio.Action) {
		fmt.Println("tick")
		delay = time.Second
		return
	}*/

	events.Serving = func(s evio.Server) (action evio.Action) {
		log.Printf("server started on port %d (loops: %d)", 5000, srv.NumLoops)
		srv = s
		return
	}

	events.Opened = func(c evio.Conn) (out []byte, opts evio.Options, action evio.Action) {
		fmt.Printf("server:%s new connect: %s\n", c.LocalAddr(), c.RemoteAddr())
		message = fmt.Sprintf("%s connected\n", c.RemoteAddr())
		//c.SetContext(&conn{})
		opts.TCPKeepAlive = 10 * time.Second // TCPKeepAlive
		for _, v := range s {
			v.Wake()
		}
		s = append(s, c)
		return
	}

	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		s := []byte("quit\r\n")
		if bytes.Equal(in, s) {
			return nil, evio.Close
		} else {
			out = in
		}

		if in == nil {
			//wake
			out = []byte(message)
		}
		return
	}

	if err := evio.Serve(events, "tcp://0.0.0.0:5000"); err != nil {
		panic(err.Error())
	}
}
