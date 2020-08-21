package main

import (
	"time"
	"fmt"
)

func main() {
	name, _ := time.Now().Zone();
	fmt.Println("current zone: ", name)
}
