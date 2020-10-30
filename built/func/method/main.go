package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (f *field) print() {
	fmt.Println(f.name)
}

func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print()
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		// go v.print()
		go (*field).print(&v)
	}
	time.Sleep(3 * time.Second)
}
