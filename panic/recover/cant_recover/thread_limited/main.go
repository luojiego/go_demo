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

//在 Ubuntu 18.04 上的输出
//在 Win10 上的输出比较随机
/*
change max threads 10000->10
start goroutine: #0
start goroutine #0 ok
start goroutine: #1
start goroutine #1 ok
start goroutine: #2
start goroutine #2 ok
start goroutine: #3
start goroutine #3 ok
start goroutine: #4
start goroutine #4 ok
start goroutine: #5
start goroutine #5 ok
start goroutine: #6
start goroutine #6 ok
start goroutine: #7
runtime: program exceeds 10-thread limit
fatal error: thread exhaustion

*/
