package _1

import "testing"

func BenchmarkBoxPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ip = p
	}
}

func BenchmarkBoxInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ix = x
	}
}

func BenchmarkBoxString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iy = y
	}
}

func BenchmarkBoxSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iz = z
	}
}
