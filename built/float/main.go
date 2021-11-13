package main

func main() {
	// var a int32 = 3
	// var b int32 = 3
	// var c int32 = 9
	// println(c / (a - b)) // 运行出错 panic: runtime error: integer divide by zero
	// println(c / 0)       // 编译出错 division by zero
	// println(float64(0) / float64(0)) // 编译出错 division by zero
	var a float64 = 0.0
	println(float64(0) / a)    // NaN
	println(float64(0.1) / a)  // +Inf
	println(float64(-0.1) / a) // -Inf
	var b float32 = 0.0
	println(float32(0) / b)    // NaN
	println(float32(0.1) / b)  // +Inf
	println(float32(-0.1) / b) // -Inf
}
