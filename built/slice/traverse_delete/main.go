package main

import "fmt"

var (
	s = []int{1, 2, 3, 4, 5, 6, 7, 8}
)

func main() {
	fmt.Println(s)
	for k, v := range s {
		if v%2 == 0 {
			s = append(s[:k], s[k+1:]...)
		}
	}
	fmt.Println(s)
}
