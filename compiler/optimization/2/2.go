package main

import (
	"bytes"
	"fmt"
	"testing"
)

// 三种情况
// bytes.Repeat 的 count 为 1，结果为 0,0,0
// count > 1 && count <= 32，结果为 0,0,1
// count > 32，结果为 0,1,1

var name = bytes.Repeat([]byte{'x'}, 32)
var m, s = make(map[string]string, 10), ""

func f() {
	s = m[string(name)]
}

func g() {
	key := string(name)
	s = m[key]
}

func h() {
	m[string(name)] = "Golang"
}

func main() {
	fmt.Println(testing.AllocsPerRun(1, f))
	fmt.Println(testing.AllocsPerRun(1, g))
	fmt.Println(testing.AllocsPerRun(1, h))
}
