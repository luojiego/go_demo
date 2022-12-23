package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(wg *sync.WaitGroup, cancel chan bool)  {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-cancel:
			return
		}
	}
}

func main() {
	cancel := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++{
		wg.Add(1)
		go Worker(&wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
