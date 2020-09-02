package main

import (
	"fmt"
	"testing"
)

// x 和 y 的长度均大于 32 时，结果为 1，3
// x 和 y 的长度均小于等于 32 时，结果为 1,1

var x = []byte{32: 'x'}
var y = []byte{32: 'y'}
var s string

func f() {
	s = ("-" + string(x) + string(y))[1:]
}

func g() {
	s = string(x) + string(y)
}

func main() {
	fmt.Println(testing.AllocsPerRun(1, f))
	fmt.Println(testing.AllocsPerRun(1, g))
}
