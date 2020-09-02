package main

import (
	"fmt"
	"testing"
)

// 多个字符串链接只需要开辟一次内存，无论链接多少次
// 测试结果为 1，3

var x, y, z, w = "Hello", "World", "Let's", "Go!"
var s string

func f() {
	s = x + y + z + w
}

func g() {
	s = x + y
	s += z
	s += w
}

func main() {
	fmt.Println(testing.AllocsPerRun(1, f))
	fmt.Println(testing.AllocsPerRun(1, g))
}
