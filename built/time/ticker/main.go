package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	t1 := time.NewTicker(time.Second * 10)
	t2 := time.NewTicker(time.Minute)

	for {
		select {
		case <-ticker.C:
			fmt.Println("ticker trigger")
		case <-t1.C:
			fmt.Println("10s trigger")
		case <-t2.C:
			fmt.Println("minute trigger")
		}
	}

}
