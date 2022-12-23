package main

import (
	"pool"
	"time"
)

func main() {
	p := pool.New(5, pool.WithPreAllocWorkers(false), pool.WithBlock(false))

	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 3)
		})

		if err != nil {
			println("task: ", i, "err: ", err)
		}
	}
	p.Free()
}
