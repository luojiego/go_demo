package main

import (
	"fmt"
	"math/rand"
)

func BubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n; i++ {
		var swap bool
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

func InsertSort(s []int) {
	n := len(s)
	for i := 1; i < n; i++ {
		v := s[i]
		j := i - 1
		for ; j >= 0; j-- {
			if s[j] > v {
				s[j+1] = s[j]
			} else {
				break
			}
		}
		s[j+1] = v
	}
}

// 排序区
// 未排序区
func SelectSort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		k := i
		// min := s[i]
		for j := i + 1; j < n; j++ {
			if s[j] < s[k] {
				k = j
			}
		}

		if k != i {
			s[i], s[k] = s[k], s[i]
		}
	}
}

const (
	size = 30
)

func makeSlice() []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(200)
	}
	return s
}

func main() {
	s := makeSlice()
	fmt.Println(s)
	BubbleSort(s)
	fmt.Println(s)

	s = makeSlice()
	fmt.Println(s)
	InsertSort(s)
	fmt.Println(s)

	s = makeSlice()
	fmt.Println(s)
	SelectSort(s)
	fmt.Println(s)
}
