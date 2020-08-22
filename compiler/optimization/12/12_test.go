package _2

import "testing"

func BenchmarkX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ix = x
	}
}

func BenchmarkY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iy = y
	}
}
