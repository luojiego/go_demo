package main

const c1 = 100

var v1 = 200

func main() {
	// const 不能取地址
	// println(&c1, c1)
	println(&v1, v1)
}
