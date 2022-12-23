package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[2:4:6]
	fmt.Printf("%#v len(slice)=%d cap(slice)=%d\n",
		slice, len(slice), cap(slice))
}
