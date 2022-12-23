package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type SumFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timedSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed (%s): %v ---\n",
				getFunctionName(f), time.Since(t))
		}(time.Now())
		return f(start, end)
	}
}

func sum1(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}

	var sum int64
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

func main() {
	s1 := timedSumFunc(sum1)
	s2 := timedSumFunc(sum2)
	fmt.Printf("%d, %d\n",
		s1(-1000, 100000000), s2(-1000, 100000000))
}
