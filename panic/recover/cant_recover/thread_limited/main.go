package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

func main() {
	nv := 10
	ov := debug.SetMaxThreads(nv)
	fmt.Printf("change max threads %d->%d\n", ov, nv)

	var wg sync.WaitGroup
	c := make(chan bool, 0)

	for i := 0; i < nv; i++ {
		fmt.Printf("start goroutine: #%v\n", i)

		wg.Add(1)
		go func() {
			c <- true
			defer wg.Done()
			runtime.LockOSThread()
			time.Sleep(10 * time.Second)
			fmt.Println("goroutine quit")
		}()
		<-c
		fmt.Printf("start goroutine #%v ok\n", i)
	}
	fmt.Println("wait for all goroutines about 10s...")
	wg.Wait()

	fmt.Println("all goroutines done")
}
