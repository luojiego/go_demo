package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	start := 0
	neg := 1
	for ; start < len(s); start++ {
		if s[start] == ' ' {
			continue
		} else {
			if s[start] == '+' {
				start += 1
			} else if s[start] == '-' {
				start += 1
				neg = -1
			}
			break
		}
	}

	num := 0
	for i := start; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
			if num*neg > math.MaxInt32 {
				return math.MaxInt32
			} else if num*neg < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			// 遇到非数字，则结束循环
			break
		}
	}
	return num * neg
}

func main() {
	fmt.Println(myAtoi("   42"))
	fmt.Println(myAtoi(" +42"))
	fmt.Println(myAtoi("-42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words 987"))
	fmt.Println(myAtoi("-91283472332"))
}
