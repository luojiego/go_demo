package main

import "fmt"

var (
	count = 0
)

func move(disk int, n, m string) {
	count++
	// ln("第" + (++m) +" 次移动 : " +" 把 "+ disks+" 号圆盘从 " + N +" ->移到->  " + M);
	fmt.Printf("第 %d  次移动：把 %d 号圆盘从 %s ->移到->%s\n", count, disk, n, m)
}

func hanoi(n int, a, b, c string) {
	if n == 1 {
		move(1, a, c)
	} else {
		hanoi(n-1, a, c, b)
		move(n, a, c)
		hanoi(n-1, b, a, c)
	}
}

func main() {
	fmt.Println("-------------------------------------------------")
	hanoi(1, "A", "B", "C")
	fmt.Println("-------------------------------------------------")
	hanoi(2, "A", "B", "C")
	fmt.Println("-------------------------------------------------")
	hanoi(3, "A", "B", "C")
}
