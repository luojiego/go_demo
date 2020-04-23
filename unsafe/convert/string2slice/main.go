package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func string2bytes1(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	var b []byte
	pBytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pBytes.Data = stringHeader.Data
	pBytes.Len = stringHeader.Len
	pBytes.Cap = stringHeader.Len

	runtime.KeepAlive(s)
	return b
}

func string2bytes2(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func main() {
	s1 := string2bytes1("Roger")
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s2 := string2bytes2("Roger")
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}
