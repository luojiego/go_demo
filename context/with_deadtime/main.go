package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) error {
	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Doing some work ", i)

		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}

func main() {
	d := time.Now().Add(6 * time.Second)
	//设置自动结束的时间
	ctx, _ := context.WithDeadline(context.Background(), d)

	//defer cancel()

	go work(ctx)

	time.Sleep(10 * time.Second)
}
