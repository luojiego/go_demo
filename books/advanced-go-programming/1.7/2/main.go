package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	panic(1)
}
/*
func MyRecover() interface{}   {
	log.Println("trace...")
	return recover()
}*/
