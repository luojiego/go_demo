package main

import "fmt"

func removeElement(nums []int, val int) int {
	l := len(nums)
	for i := 0; i < l; {
		if nums[i] == val {
			nums[i] = nums[l-1]
			l--
		} else {
			i++
		}
	}
	return l
}

func main() {
	// nums := []int{3,2,2,3}
	// val := 3

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	n := removeElement(nums, val)
	for i := 0; i < n; i++ {
		fmt.Println(nums[i])
	}
}
