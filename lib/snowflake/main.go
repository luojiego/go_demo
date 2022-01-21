package main

import (
	"fmt"

	"github.com/luojiego/snowflake"
)

func main() {
	node, err := snowflake.NewNode(3)
	if err != nil {
		panic(err)
	}
	uuid := node.Generate()
	fmt.Printf("%+v\n", uuid)
}
