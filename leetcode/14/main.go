package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var ret []uint8
	for i := 0; ; i++ {
		var c uint8
		f := false
		for k, v := range strs {
			if len(v) <= i {
				break
			} else {
				if k == 0 {
					c = v[i]
				} else {
					if c != v[i] {
						break
					}
					if k == len(strs)-1 {
						f = true
					}
				}
			}
		}

		if !f {
			break
		} else {
			ret = append(ret, c)
		}
	}

	return string(ret)
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
}
