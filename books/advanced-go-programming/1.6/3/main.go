package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Producer(num int, out chan<- int) {
	for i := 1; ; i++ {
		out <- i * num
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

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
