package main

import (
	"fmt"
	"reflect"
)

type S struct {
	Num int
	Str string
	S   []int
}

func main() {
	a := S{10, "hello", []int{1}}
	b := S{10, "hello", []int{1}}
	fmt.Println(&a == &b)
	fmt.Println(reflect.DeepEqual(a, b))
}
