package main

import (
	"fmt"
	"io"
	"unsafe"
)

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s lauguage\n", p.language)
}

func (p Gopher) debug() {
	fmt.Printf("I am debuging %s lauguage\n", p.language)
}

func main() {
	{
		x := 300
		var any interface{} = x
		fmt.Println(any)

		g := Gopher{"Go"}
		var c coder = g
		fmt.Println(c)
	}

	{
		var c coder
		fmt.Println(c == nil)
		fmt.Printf("c: %T, %v\n", c, c)

		var g *Gopher
		fmt.Println(g == nil)

		c = g
		fmt.Println(c == nil)
		fmt.Printf("c: %T, %v\n", c, c)
	}

	{
		type iface struct {
			itab, data uintptr
		}

		var a interface{} = nil
		var b interface{} = (*int)(nil)
		x := 5
		var c interface{} = (*int)(&x)

		ia := *(*iface)(unsafe.Pointer(&a))
		ib := *(*iface)(unsafe.Pointer(&b))
		ic := *(*iface)(unsafe.Pointer(&c))

		fmt.Println(ia, ib, ic)
		fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
	}

	{
		var _ io.Writer = (*myWriter)(nil)
		var _ io.Writer = myWriter{}
	}
}

type myWriter struct {
}

func (w myWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
