package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	v, ok := m.Load("123")
	if !ok {
		fmt.Println("can't find 123")
	}
	fmt.Println(v)
}
