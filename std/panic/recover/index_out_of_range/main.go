package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover: ", r)
		}
	}()

	s := []int{1, 2}
	fmt.Println(s[2])
}
