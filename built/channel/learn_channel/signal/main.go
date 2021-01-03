package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(time.Now().String())
		}
	}()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	doCleanup()

	fmt.Println("优雅退出")
}

func doCleanup() {
	fmt.Println("总有刁民想害联")
}
