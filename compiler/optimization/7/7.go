package clear

const N = 1024 * 100

var a [N]int

func init() {
	for i := range a {
		a[i] = 1024
	}
}

// go test -bench=.

func clearArray() {
	for i := range a {
		a[i] = 0
	}
}

func clearSlice() {
	s := a[:]
	for i := range s {
		s[i] = 0
	}
}

func clearArrayPtr() {
	for i := range &a {
		a[i] = 0
	}
}
