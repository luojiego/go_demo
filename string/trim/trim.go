package main

import (
	"fmt"
	"strings"
)

var (
	filename = "hello.exe"
)

// 我简直就是个大傻子，用 TrimLeft 想去掉 .exe，弄的自己怀疑人生
func main() {
	fmt.Println(strings.TrimRight(filename, ".exe"))
}
