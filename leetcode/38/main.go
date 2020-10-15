package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	s := "1"
	if n == 1 {
		return s
	}

	for i := 2; i <= n; i++ {
		count := 1
		c := s[0]
		tmp := ""
		for j := 1; j < len(s); j++ {
			if c == s[j] {
				count++
			} else {
				tmp += strconv.Itoa(count) + string(c)
				c = s[j]
				count = 1
			}
		}
		s = tmp + strconv.Itoa(count) + string(c)
	}
	return s
}

func main() {
	// c := '1'
	// fmt.Println(string(c))
	for i := 1; i <= 10; i++ {
		fmt.Println(i, ":", countAndSay(i))
	}
}
