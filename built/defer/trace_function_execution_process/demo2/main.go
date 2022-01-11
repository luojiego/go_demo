package main

import (
	"fmt"
	"runtime"
)

func Trace() func() {
	// 自动获取函数名称
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	fmt.Printf("enter: [%s]\n", name)
	return func() {
		fmt.Printf("exit: [%s]\n", name)
	}
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}

func main() {
	defer Trace()()
	foo()
}

// 输出结果
/*
enter: main
enter: foo
enter: bar
exit: bar
exit: foo
exit: main
*/
