package main

import (
	"fmt"

	"github.com/sqids/sqids-go"
)

func main() {
	fmt.Println("vim-go")
	s, _ := sqids.New(sqids.Options{
		Alphabet: "TDY8UAV2RG75MQ3LKZBSW9HJCX6PNF4E",
	})
	numbers := s.Decode("CY7B5DTJ")
	fmt.Println(numbers) // [1 6]
}
