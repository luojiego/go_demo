package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name    string
	Age     int
	Address string
}

type B struct {
	A
	Name string
}

func main() {
	b := B{
		A: A{
			Name:    "Roger",
			Age:     30,
			Address: "ABC",
		},
		Name: "罗杰",
	}
	r, _ := json.Marshal(b)
	fmt.Println(string(r))

	s := `{"Age":30,"Address":"ABC","Name":"罗杰"}`
	b1 := &B{}
	json.Unmarshal([]byte(s), b1)
	fmt.Println(b1)
}
