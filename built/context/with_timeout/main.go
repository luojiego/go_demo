package main

import (
	"fmt"
	"sync"
	"time"

	"context" //博客中的 context import 的是 "golang.org/x/net/context"
)

//https://juejin.im/post/5a6873fef265da3e317e55b6
var (
	wg sync.WaitGroup
)

func work(ctx context.Context) error {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Doing some work ", i)

		// we received the signal of cancelation in this channel
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel() //goroutine 执行结束 立刻释放资源

	fmt.Println("Hey, I'm going to do some work")

	wg.Add(1)
	go work(ctx)
	wg.Wait()

	fmt.Println("Finished. I'm going home")
}
