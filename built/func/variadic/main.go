package main

import "fmt"

func dump(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	// s := []string{"Tony", "John", "Jim"}
	// dump(s...) // cannot use s (type []string) as type []interface {} in argument to dump
	s := []interface{}{"Tony", "John", "Jim"}
	dump(s...)
}
