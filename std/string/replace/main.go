package main

import (
	"fmt"
	"strings"
)

//使用 [:4] 这样的方式取得 str 的前4个字符赋值给 s，然后s使用strings的replace方法, 会不会修改原字符str

const (
	name = `luo"jie"`
)

func main() {
	str := "ABCDEFGHIJKLMN"

	s := str[:4]

	s = strings.ReplaceAll(s, "A", "a")
	s = strings.ReplaceAll(s, "B", "b")

	name1 := strings.ReplaceAll(name, "\"", "\"\"")

	fmt.Println(s, str)

	fmt.Printf("name: %s\n", name1)
}
