package main

import "fmt"

type S struct {
	m map[string]interface{}
}

func main() {
	var m map[string]interface{}
	if m == nil {
		println("m is nil")
	}

	var s S
	if s.m == nil {
		println("s.m is nil")
	}

	// Go 夜读群里面有人说用了 new 之后，map 会被初始化
	// 结果并不是他说的这么回事
	s1 := new(S)
	if s1.m == nil {
		println("s1.m is nil")
	} else {
		println("s1.m is not nil")
	}

	// 但是对一个 nil 的 map 可以取值
	if _, ok := s1.m["2"]; !ok {
		println("can't find 2")
	}

	// 直接打印也是可以的，会输出 nil
	fmt.Println(s1.m["1"])
}
