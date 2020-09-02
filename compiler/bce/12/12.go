package _2

func f1(s []int) {
	_ = s[0]
	_ = s[1]
	_ = s[2]
}

func f2(s []int) {
	_ = s[2]
	_ = s[1]
	_ = s[0]
}

func f3(s []int) {
	if len(s) > 2 {
		_, _, _ = s[0], s[1], s[2]
	}
}

func f4(a [5]int) {
	_ = a[4]
}

func f5(s []int, index int) {
	if index >= 0 && index < len(s) {
		_ = s[index]
		_ = s[index:len(s)]
	}
}

func f6(s []int, index int) {
	_ = s[index]
	_ = s[index]
}

func f7(s []int) {
	for i := range s {
		_ = s[i]
		_ = s[i:len(s)]
		_ = s[:i+1]
	}
}

func f8(s []int) {
	for i := 0; i < len(s); i++ {
		_ = s[i]
		_ = s[i:]
		_ = s[:i+1]
	}
}

func f9(s []int) {
	for i := len(s) - 1; i >= 0; i-- {
		_ = s[i]
		_ = s[i:]
	}
}

func fd(is []int, bs []byte) {
	if len(is) >= 256 {
		for _, n := range bs {
			_ = is[n]
		}
	}
}

func fd2(is []int, bs []byte) {
	if len(is) >= 256 {
		is = is[:256]
		for _, n := range bs {
			_ = is[n]
		}
	}
}

func fe(isa []int, isb []int) {
	if len(isa) > 0xFFF {
		for _, n := range isb {
			_ = isa[n&0xFFF]
		}
	}
}

func fe2(isa []int, isb []int) {
	if len(isa) > 0xFFF {
		isa = isa[:0xFFF+1]
		for _, n := range isb {
			_ = isa[n&0xFFF]
		}
	}
}

func fp(x, y string) int {
	if len(x) > len(y) {
		x, y = y, x
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return i
		}
	}

	return len(x)
}

func fq(x, y string) int {
	if len(x) > len(y) {
		x, y = y, x
	}

	if len(x) <= len(y) {
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] {
				return i
			}
		}
	}
	return len(x)
}
