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
	fmt.Printf("s1: %s, len:%d, cap:%d\n", string(s1), len(s1), cap(s1))
	s1 = append(s1, " Golang"...)
	fmt.Printf("s1: %s, len:%d, cap:%d\n", string(s1), len(s1), cap(s1))
	s2 := string2bytes2("Roger")
	fmt.Printf("s2: %s, len:%d, cap:%d\n", string(s2), len(s2), cap(s2))
	s2 = append(s2, " Golang"...)
	fmt.Printf("s2: %s, len:%d, cap:%d\n", string(s2), len(s2), cap(s2))
}
