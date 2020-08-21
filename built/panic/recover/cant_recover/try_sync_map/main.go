package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := sync.Map{}

	f := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover: ", r)
			}
		}()
		for {
			m.Store("name", "Roger")
		}
	}

	go f()
	go f()

	time.Sleep(1 * time.Second)
}
