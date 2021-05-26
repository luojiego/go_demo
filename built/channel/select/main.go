package main

import (
	"fmt"
	"time"
)

func unBufferChannel() {
	println("unBufferChannel")
	var ch = make(chan int)
	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			println(v)
		}
	}
}

func bufferChannel() {
	println("bufferChannel")
	var ch = make(chan int, 10)
	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			println(v)
		}
	}
}

type TestData struct {
	Name string
	Age  int
}

func main() {
	var ch = make(chan *TestData, 100)
	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
			default:
				// 使用 select 真的非常方便，可以检测接收的数据
				// fmt.Println("error")
				// 如果没有消息，则等 1s，避免对 CPU 造成过大压力
				time.Sleep(1 * time.Second)
			}
		}
	}()
	// unBufferChannel()
	// bufferChannel()

	for i := 0; i < 20; i++ {
		ch <- &TestData{
			Name: "罗杰",
			Age:  30 + i,
		}
	}
	select {}
}
