package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Once struct {
	m sync.Mutex
	done uint32
}

func (o *Once) Do (f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func main() {
	o := Once{}
	o.Do(func() {
		fmt.Println("Once do")
	})
}
