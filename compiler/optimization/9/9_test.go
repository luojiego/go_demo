package struct_fileds

import "testing"

func BenchmarkRange1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range ss1 {
			x1 = v.a
		}
	}
}

func BenchmarkRange2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range ss2 {
			x2 = v.a
		}
	}
}

func BenchmarkRange3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range ss3 {
			x3 = v.a
		}
	}
}

func BenchmarkRange4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range ss4 {
			x4 = v.a
		}
	}
}

func BenchmarkRange5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range ss5 {
			x5 = v.a
		}
	}
}

func BenchmarkRange6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range ss6 {
			x6 = v.a
		}
	}
}
