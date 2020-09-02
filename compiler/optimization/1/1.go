package main

import (
	"fmt"
	"strings"
	"testing"
)

// 尝试发现 32 位字符串以上时，g 函数会有内存分配发生，但是无论有多少长度的字符串
// f 函数并不会有内存分配
var (
	str = strings.Repeat("go", 16)
)

func f() {
	for range []byte(str) {

	}
}

func g() {
	b := []byte(str)
	for range b {

	}
}

func main() {
	fmt.Println(testing.AllocsPerRun(1, f))
	fmt.Println(testing.AllocsPerRun(1, g))
}
