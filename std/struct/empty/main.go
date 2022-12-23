package main

import "fmt"

//注意：println 和 fmt.Println 的目标输出是不一致，很有可能下面的两条输出顺序会不是固定的
//所以我在前面加上标识
func main() {
	a := new(struct{})
	b := new(struct{})
	println("a: ", a, b, a == b)
	c := new(struct{})
	d := new(struct{})
	fmt.Println("c: ", c, d, c == d)
}

//output
/*
a:  0xc00010ff37 0xc00010ff37 false
c:  &{} &{} true
*/
//使用以下 命令查看 逃逸分析
// go run -gcflags="-m" main.go

//为什么 a == b 为 false
//使用 GOSSAFUNC=main go build main.go 来生成 saa.html 来分析
//可以看到 代码优化阶段（opt）直接被编译器作为常量优化掉了

//使用以下命令来停止编译器优化
//go run -gcflags="-N -l" main.go
