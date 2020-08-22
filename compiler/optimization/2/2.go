package main

import (
	"bytes"
	"fmt"
	"testing"
)

var name = bytes.Repeat([]byte{'x'}, 33)
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
