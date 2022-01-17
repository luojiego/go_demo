package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	Age int `json:"age"`
	Number int `json:"number"`
}

func (t Test) MarshalJSON() ([]byte, error) {
	return []byte("{}"), nil//fmt.Errorf("what?")
}

func main() {
	d := struct {
		Test
		Name string `json:"name"`
	}{
		Test{
			Age: 30, Number: 1,
		},"LuoJie",
	}
	result, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}
