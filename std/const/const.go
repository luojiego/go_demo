package main

import "math"

const c1 = 100

var v1 = 200

const (
	a, b = 10, 20
	// arr  = [3]int{1, 2, 3}
	t2 = 2147483647
)

const (
	Apple, Banana     = 11, 22
	Strawberry, Grape // 11, 22
	Pear, Watermelon  // 11, 22
)

const (
	Apple1      = 1
	Strawberry1 // 1
)

const (
	mutexLocked           = 1 << iota // 1 << 0 = 1
	mutexWoken                        // 1 << 1 = 2
	mutexStarving                     // 1 << 2 = 4
	mutexWaiterShift      = iota      // 3
	starvationThresholdNs = 1e6       // 1e6
)

func main() {
	println(mutexStarving)
	println(mutexWaiterShift)
	// const 不能取地址
	// println(&c1, c1)
	println(&v1, v1)

	var c int64 = 300
	var d uint32 = 200
	var f float64 = 100.0
	var g int8 = 126
	println(c + a)
	println(d + b)
	println(f + a)
	println(g + b)
	var t1 int32 = int32(math.Pow(2, 30))
	// var t2 int32 = int32(math.Pow(2, 30))
	println(t1)
	println(t1 + t2)
}
