package clear

import (
	"testing"
)

func BenchmarkClearArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clearArray()
	}
}

func BenchmarkClearSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clearSlice()
	}
}

func BenchmarkClearArrayPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clearArrayPtr()
	}
}
