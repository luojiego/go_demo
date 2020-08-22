package main

import (
	"fmt"
	"strings"
	"testing"
)

// 若 GoGoGo 的长度大于 32 时，结果为 0，1
// 若 GoGoGo 的长度小于等于 32 时，结果为 0，0

var GoGoGo = strings.Repeat("Go", 17)

func f() {
	_ = len([]rune(GoGoGo))
}

func g() {
	_ = len([]byte(GoGoGo))
}

func main() {
	fmt.Println(testing.AllocsPerRun(1, f))
	fmt.Println(testing.AllocsPerRun(1, g))
}
