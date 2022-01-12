package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 16, 32)
	len := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(len)

	cap := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(cap)
}
