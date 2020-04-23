package main

import (
	"fmt"
	"unsafe"
)

func slice2bytes(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}

func main() {
	s := []byte{'R', 'o', 'g', 'e', 'r'}
	fmt.Println(s)
	fmt.Println("len(string):", len(s))
	fmt.Println("cap(string):", cap(s))

	r := slice2bytes(s)
	fmt.Println("string:", r)
	fmt.Println("len(string):", len(r))
}
