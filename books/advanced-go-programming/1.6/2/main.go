package main

import (
	"fmt"
	"time"
)

func Producer(factor int, out chan<- int) {
	for i := 1; ; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)

	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)

	time.Sleep(3 * time.Second)
}
