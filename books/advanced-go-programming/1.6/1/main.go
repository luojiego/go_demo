package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("hello world!")
			wg.Done()
		}()
	}

	wg.Wait()
}
