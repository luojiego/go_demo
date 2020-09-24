package main

import "fmt"

func main() {
	m := make(map[string]map[string]interface{})
	m["一"]["1"] = 111
	fmt.Println(m)
}
