package main

import "fmt"

func double(x *int) {
	*x += *x
	x = nil
}

func main() {
	a := 3
	double(&a)
	fmt.Println(a)

	p := &a
	double(p)
	fmt.Println(a, p == nil)
}
