package main

import (
	"context"
	"fmt"
	"time"
)

//黑人？？ 不太明白
func main() {
	d := time.Now().Add(6 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	go func(ctx context.Context) {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("oversleep")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}
	}(ctx)

	time.Sleep(10 * time.Second)
}
