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

//对于cancel 好多人 都有误解
//1. context 设计的目的在于 goroutine 结束之后尽快释放掉资源
//		1.1 显然下面的代码是有问题的 6s goroutine 便会执行结束，但是 10s main函数退出时才会执行cancel函数
//		实际使用中，可以借助 sync.WaitGroup 或者其它方式来实现
//2. 部分文章说明是错误的，说是 goroutine 结束之后才能调用cancel
//		2.1 实际测试时，我取掉了 defer cancel() 中的 defer 程序会直接打印 "Cancel the context  0"

func main() {
	d := time.Now().Add(6 * time.Second)
	//设置自动结束的时间
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	go work(ctx)

	time.Sleep(10 * time.Second)
}
