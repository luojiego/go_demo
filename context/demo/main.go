package main

import (
	"context"
	"fmt"
	"time"
)

//https://qcrao.com/2019/06/12/dive-into-go-context/
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("monitor quit, stop...")
				return
			default:
				fmt.Println("goroutines monitoring")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("quit...")
	cancel()
	time.Sleep(3 * time.Second)
}
