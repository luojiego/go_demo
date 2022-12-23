package main

import (
	"fmt"
)

// 编程实现“求一个数的平方根”
// 牛顿弦切法
func sqrt(n float64) float64 {
	if n < 0 {
		return -1
	}

	root := n
	for root*root-n > 1e-6 {
		root = (n + root*root) / 2 / root
	}

	// for math.Abs(n-root*root) >= 1e-6 {
	// 	root = (n/root + root) / 2.0
	// }
	return root
}

func main() {
	fmt.Println(sqrt(8))
	fmt.Println(sqrt(9))
}
