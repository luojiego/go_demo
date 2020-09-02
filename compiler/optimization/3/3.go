package main

import (
	"fmt"
	"testing"
)

// demo 中的值是 1023
// 实际测试结果
// x 和 y 的长度大于 32 时，结果为 0，2
// x 和 y 的长度小于等于 32 时，结果为 0，0

var x = []byte{32: 'x'}
var y = []byte{32: 'y'}
var b bool

func f() {
	b = string(x) != string(y)
}

func g() {
	sx, sy := string(x), string(y)
	b = sx == sy
}

func main() {
	fmt.Println(len(x), len(y))
	fmt.Println(testing.AllocsPerRun(1, f))
	fmt.Println(testing.AllocsPerRun(1, g))
}
