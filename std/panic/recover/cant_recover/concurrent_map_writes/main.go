package main

import (
	"fmt"
	"time"
)

//fatal error: concurrent map writes

//可以使用 sync.Map 来解决
func main() {
	m := map[string]int{}
	p := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover: ", r)
			}
		}()

		for {
			m["t"] = 0
		}
	}
	go p()
	go p()
	time.Sleep(1 * time.Second)
}
