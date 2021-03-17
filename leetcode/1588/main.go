package main

import "fmt"

func sumOddLengthSubArrays(arr []int) int {
	var r int
	for i := 0; i < len(arr); {
		for j := 0; j < len(arr); j++ {
			// fmt.Printf("i[%d]====j[%d]\n", i, j)
			if i + j >= len(arr) {
				break
			}
			for k:=0; k<=i; k++ {
				r += arr[j+k]
				// fmt.Println("---->", arr[j+k])
			}
		}
		i += 2
		// fmt.Println(r)
	}
	return r
}

func main() {
	fmt.Println(sumOddLengthSubArrays([]int{1,4,2,5,3}))
	fmt.Println(sumOddLengthSubArrays([]int{1,2}))
	fmt.Println(sumOddLengthSubArrays([]int{10,11,12}))
}
