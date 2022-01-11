package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		if r := MyRecover(); r != nil {
			fmt.Println(r)
		}
	}()

	panic(1)
}

func MyRecover() interface{}   {
	log.Println("trace...")
	return recover()
}
