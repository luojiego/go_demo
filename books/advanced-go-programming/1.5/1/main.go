package main

import (
	"fmt"
	"sync"
)

var Total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <=100; i++ {
		Total.Lock()
		Total.value += i
		Total.Unlock()
	}
}

func main() {
	fmt.Println(Total.value)
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(Total.value)
}
