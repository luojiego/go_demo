package main

import "fmt"

const (
//go 中的无法定义 const 的数组
//arr = [...]int{1, 2, 3, 4, 5, 6, 7}
)

var arr = [7]int{1, 2, 3, 4, 5, 6, 7}

func main() {
	s1 := arr[1:5]
	s2 := arr[2:6]

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	s1[1] = 200

	fmt.Println(arr, len(arr))
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
}
