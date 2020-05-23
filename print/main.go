package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (d Data) String() string {
	d.Age = 18
	r, _ := json.Marshal(d)
	return string(r)
}

func main() {

	d := Data{
		Name: "罗杰",
		Age:  80,
	}
	fmt.Println(d)
}
