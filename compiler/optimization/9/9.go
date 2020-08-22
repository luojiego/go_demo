package struct_fileds

type s1 struct{ a int }
type s2 struct{ a, b int }
type s3 struct{ a, b, c int }
type s4 struct{ a, b, c, d int }
type s5 struct{ a, b, c, d, e int }
type s6 struct{ a, b, c, d, e, f int }

var ss1, ss2, ss3, ss4, ss5, ss6 = make([]s1, 1000), make([]s2, 1000),
	make([]s3, 1000), make([]s4, 1000),
	make([]s5, 1000), make([]s6, 1000)

var x1, x2, x3, x4, x5, x6 int
