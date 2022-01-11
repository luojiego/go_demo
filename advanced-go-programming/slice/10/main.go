package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s1 []int
	var s2 = []int{}
	var s3 = make([]int, 0)
	var s4 = *new(int)

	var a1 = *(*[3]int)(unsafe.Pointer(&s1))
	var a2 = *(*[3]int)(unsafe.Pointer(&s2))
	var a3 = *(*[3]int)(unsafe.Pointer(&s3))
	var a4 = *(*[3]int)(unsafe.Pointer(&s4))

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	s2 = append(s2, 1)
	fmt.Println(s2)
	fmt.Println(s3)
}
