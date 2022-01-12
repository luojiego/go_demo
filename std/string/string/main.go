package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s = "中国人"
	fmt.Printf("the %s length:  = %d\n", s, len(s)) // len = 9
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x", s[i])
	}
	fmt.Printf("the character count s is: %d\n", utf8.RuneCountInString(s)) // 3
	for _, c := range s {
		fmt.Printf("0x%x ", c)
	}
	fmt.Printf("\n")
}
