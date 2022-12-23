package main

import (
	"fmt"
	"os"
	"strings"
)

type MyString string

func (m MyString) String() string {
	return strings.ToUpper(string(m))
}

func main() {
	fmt.Fprintln(os.Stdout, MyString("my name is luojie"))
}
