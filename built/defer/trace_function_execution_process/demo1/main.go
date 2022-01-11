package main

func Trace(name string) func() {
	println("enter:", name)
	return func() {
		println("exit:", name)
	}
}

func foo() {
	defer Trace("foo")()
	bar()
}

func bar() {
	defer Trace("bar")()
}

func main() {
	defer Trace("main")()
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
