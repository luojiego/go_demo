package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m sync.Mutex
)

func test(n int) {
	m.Lock()
	defer m.Unlock()
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(n)
}

func main() {
	test(5)
	test(1)
}
