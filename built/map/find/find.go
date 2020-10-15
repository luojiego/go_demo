package main

import "fmt"

func main() {

	m := make(map[int]interface{})
	for i := 0; i < 10; i++ {
		if m[i] != nil {
			fmt.Printf("%d is not nil\n", i)
		}
	}

	for k, v := range m {
		fmt.Printf("key: %d v:%v\n", k, v)
	}
}
