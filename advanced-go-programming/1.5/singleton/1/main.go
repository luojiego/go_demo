package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Singleton struct {
}

var (
	instance    *Singleton
	initialized uint32
	mu          sync.Mutex
)

func Instance() *Singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &Singleton{}
	}
	return instance
}

func main() {
	i := Instance()
	fmt.Printf("%v\n", i)
}
