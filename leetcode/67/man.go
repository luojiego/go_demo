package main

import "fmt"

func addBinary(a string, b string) string {
	// 从最后一个位置开始加
	//    001 -> 3
	// 101000 -> 6
	// 101001
	la, lb, i, ret, f := len(a), len(b), 0, "", uint8(0)
	for i = 1; ; i++ {
		if i > lb || i > la {
			break
		}
		c := a[la-i] - '0' + b[lb-i] - '0' + f
		if c >= 2 {
			f = 1
			ret = string(c%2+'0') + ret
		} else {
			ret = string(c+'0') + ret
			f = 0
		}
	}

	// fmt.Println(i)
	for ; i <= lb; i++ {
		c := b[lb-i] - '0' + f
		if c >= 2 {
			f = 1
			ret = string(c%2+'0') + ret
		} else {
			ret = string(c+'0') + ret
			f = 0
		}
	}

	for ; i <= la; i++ {
		c := a[la-i] - '0' + f
		if c == 2 {
			f = 1
			ret = string(c%2+'0') + ret
		} else {
			ret = string(c+'0') + ret
			f = 0
		}
	}

	if f == 1 {
		ret = "1" + ret
	}

	return ret
}

func isRightChar(c uint8) bool {
	if (c >= '0' && c <= '9') ||
		(c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func isEqual(a, b uint8) bool {
	if a == b {
		return true
	}

	if a > '9' && b > '9' {
		if a-b == 'a'-'A' || b-a == 'a'-'A' {
			return true
		}
	}

	return false
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isRightChar(s[i]) {
			i++
			continue
		}

		if !isRightChar(s[j]) {
			j--
			continue
		}

		if !isEqual(s[i], s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(addBinary("101", "011"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("1111", "1111"))
	fmt.Println(addBinary("10100", "011"))
	fmt.Println(addBinary("10100101", "11011"))
	fmt.Println(isPalindrome("Aa"))
	fmt.Println("0P -> ", isPalindrome("0P"))
	fmt.Println(isPalindrome("1A a1"))
	fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
