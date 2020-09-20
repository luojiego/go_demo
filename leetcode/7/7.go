package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	max := int(math.Pow(2, 31) - 1)
	min := int(math.Pow(-2, 31))

	var s []int
	num := x

	if x < 0 {
		num = -x
	}

	for {
		if num < 10 {
			s = append(s, num)
			break
		}
		s = append(s, num%10)
		num /= 10

	}

	// 3 2 1
	// fmt.Println(s)

	ret := 0
	for i := 0; i < len(s); i++ {
		ret = ret*10 + s[i]
	}

	if x < 0 {
		ret = -ret
	}

	if ret < min {
		return 0
	}

	if ret > max {
		return 0
	}

	return ret
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(-120))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
	fmt.Println(reverse(1))
	fmt.Println(int(math.Pow(2, 31) - 1))
	fmt.Println(int(math.Pow(-2, 31)))
	//fmt.Println(1534236469)
	//fmt.Println(1534236469 > int(math.Pow(2, 31)-1))
	//fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-1563847412))
}
