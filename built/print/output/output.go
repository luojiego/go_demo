package main

import (
	"fmt"
	"log"
)

func main() {
	array := []int{100, 200, 300, 400, 500, 600}
	for _, v := range array {
		fmt.Printf("%d\n", v)
		log.Printf("%d\n", v)
	}
}
