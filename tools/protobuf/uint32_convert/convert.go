package main

import (
	"fmt"
	"math"
)

func main() {
	var count uint32
	fmt.Printf("count: %d\n", count)
	fmt.Printf("count - 1: %d\n", count-1)
	// num := math.Pow(2, float64(count-1))
	fmt.Printf("math.Pow(2, float64(count-1)): %f\n", math.Pow(2, float64(count-1)))
	fmt.Printf("int(math.Pow(2, float64(count-1))): %d\n", int(math.Pow(2, float64(count-1))))
	fmt.Printf("uint32(math.Pow(2, float64(count-1))): %d\n", uint32(math.Pow(2, float64(count-1))))
	fmt.Printf("int32(math.Pow(2, float64(count-1))): %d\n", int32(math.Pow(2, float64(count-1))))
	fmt.Printf("uint64(math.Pow(2, float64(count-1))): %d\n", uint64(math.Pow(2, float64(count-1))))

	fmt.Printf("int64(pow(2,1023): %d\n", int64(math.Pow(2, 1023)))
	fmt.Printf("uint64(pow(2,1023): %d\n", uint64(math.Pow(2, 1023)))
}
