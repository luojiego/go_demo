package _0

import "testing"

func Benchmark_BoxPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ip = p
	}
}

func Benchmark_PointerAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p = ip.(*[100]int)
	}
}

func Benchmark_PointerAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p = p2
	}
}
