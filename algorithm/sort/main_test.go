package main

import (
	"math/rand"
	"testing"
)

var testSlice = make([]int, 20000)

func init() {
	for i := 0; i < 20000; i++ {
		testSlice[i] = rand.Intn(20000)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testSlice)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort(testSlice)
	}
}
