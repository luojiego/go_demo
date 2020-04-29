package main

import "fmt"

//recover 只有在 defer 中才能使用

func main() {
	f := func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}

	defer func() {
		f()
	}()

	panic("ok")
}
