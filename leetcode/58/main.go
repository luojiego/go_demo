package main

func lengthOfLastWord(s string) int {
	ret := 0
	start := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if start {
				return ret
			}
		} else {
			ret++
			if start == false {
				start = true
			}
		}
	}
	return ret
}

func main() {
	println(lengthOfLastWord("Hello World  "))
	println(lengthOfLastWord(""))
	println(lengthOfLastWord("abc HI"))
}
