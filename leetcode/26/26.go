package main

import "fmt"

func removeDuplicates(nums []int) int {
	// if len(nums) == 0 {
	//	return 0
	// }
	l := 0
	for i := 1; i < len(nums); i++ {
		if nums[l] != nums[i] {
			nums[l+1] = nums[i]
			l++
		}
	}
	return l + 1
}

func main() {
	// nums := []int{1,1,2}
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(nums)
	for i := 0; i < n; i++ {
		fmt.Println(nums[i])
	}
}
