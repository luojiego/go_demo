package main

import (
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m = make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		m[i] = i
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 100; j++ {
				n := rand.Intn(100)
				m[n] = n
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
