package main

import "fmt"

func main() {
	m := make(map[string]int, 0)
	m["一"] = 1
	m["二"] = 2
	m["三"] = 3
	fmt.Println(m)

	m = make(map[string]int, 0)
	m["壹"] = 1
	m["贰"] = 2
	m["叁"] = 3
	fmt.Println(m)
}
