package main

import (
	"fmt"
	"time"
	"runtime"
)

func deadloop() {
	for {
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go deadloop() 
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}
