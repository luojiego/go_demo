package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover: ", r)
		}
	}()

	c := make(chan int)
	close(c)
	c <- 1

}
