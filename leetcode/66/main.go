package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i] + 1
		if n > 9 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
				return digits
			}
		} else {
			digits[i] = n
			return digits
		}
	}

	return nil
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{}))
	fmt.Println(plusOne(nil))
}
