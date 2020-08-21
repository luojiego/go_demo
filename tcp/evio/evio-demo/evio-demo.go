package main

import (
	"github.com/tidwall/evio"
	"log"
)

func main() {
	var events evio.Events
	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {

		if string(in) == "quit\r\n" {
			return nil, evio.Close
		} else {
			out = in
		}
		return
	}
	events.Opened = func(c evio.Conn) (out []byte, opts evio.Options, action evio.Action) {
		out = []byte("欢迎来到 1024\r\n")
		return
	}
	events.PreWrite = func() {
		log.Println("what?")
	}
	if err := evio.Serve(events, "tcp://localhost:5000"); err != nil {
		panic(err.Error())
	}
}
