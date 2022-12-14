package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // 此处对于 i 的使用会产生 trace
			// 解决方法就是将 i 以参数的形式传递进来
			wg.Done()
		}()
	}

	wg.Wait()
}
