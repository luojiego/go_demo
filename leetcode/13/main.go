package main

import "fmt"

// 字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

func romanToInt(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i < len(s)-1 {
				if s[i+1] == 'V' {
					num += 4
					i++
				} else if s[i+1] == 'X' {
					num += 9
					i++
				} else {
					num += 1
				}
			} else {
				num += 1
			}
		case 'V':
			num += 5
		case 'X':
			if i < len(s)-1 {
				if s[i+1] == 'L' {
					num += 40
					i++
				} else if s[i+1] == 'C' {
					num += 90
					i++
				} else {
					num += 10
				}
			} else {
				num += 10
			}
		case 'L':
			num += 50
		case 'C':
			if i < len(s)-1 {
				if s[i+1] == 'D' {
					num += 400
					i++
				} else if s[i+1] == 'M' {
					num += 900
					i++
				} else {
					num += 100
				}
			} else {
				num += 100
			}
		case 'D':
			num += 500
		case 'M':
			num += 1000
		default:
			num += 0
		}
	}

	return num
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
