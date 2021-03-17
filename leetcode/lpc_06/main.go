package main

import "fmt"

func minCounts(coins []int) int {
	c := 0
	for _, v := range coins {
		if v < 3 {
			c++
		} else {
			if v % 2 == 0 {
				c += v >> 1
			} else {
				c += v >> 1 + 1
			}
		}

		fmt.Println(c)
	}
	return c
}

func main() {
	fmt.Println(4 >> 1)
	fmt.Println(minCounts([]int{4,2,1}))
}
