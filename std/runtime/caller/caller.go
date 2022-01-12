package main

import (
	"fmt"
	"runtime"
)

func func1() {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		fmt.Println(file, line)
	}
}

func func2() {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println(file, line)
	}
	func1()
}

func main() {
	func2()
}
