package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		flag := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				flag = false
			}
		}
		if flag {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaaa", "bba"))
	fmt.Println(strStr("a", "a"))
}
