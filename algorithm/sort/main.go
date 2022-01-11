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

/*
// 归并排序
func mergeSort(s []int) {
	mergeSortC(s, 0, len(s))
}

var count = 0

func mergeSortC(s []int, min, max int) {
	count++
	if count >= 20 {
		os.Exit(1)
	}
	if min >= max {
		return
	}
	fmt.Printf("min: %d max: %d s:%+v\n", min, max, s[min:max])

	middle := (min + max) / 2
	mergeSortC(s, min, middle)
	if (min+max)%2 == 0 {
		mergeSortC(s, middle, max)
	} else {
		mergeSortC(s, middle, max)
	}

	merge(s[min:max], s[min:middle], s[middle:max])
	fmt.Printf("min: %d max: %d s: %+v\n", min, max, s[min:max])
}
*/

func mergeSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	first := mergeSort(s[:len(s)/2])
	second := mergeSort(s[len(s)/2:])
	return merge(first, second)
}

func merge(first, second []int) []int {
	tmp := make([]int, 0, len(first)+len(second))
	i, j := 0, 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			tmp = append(tmp, first[i])
			i++
		} else {
			tmp = append(tmp, second[j])
			j++
		}
	}
	if i < len(first) {
		tmp = append(tmp, first[i:]...)
	}
	if j < len(second) {
		tmp = append(tmp, second[j:]...)
	}
	return tmp
}

// 快速排序

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

func isSorted(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

func main() {
	// s := makeSlice()
	// fmt.Println(s)
	// BubbleSort(s)
	// fmt.Println(s)

	// s = makeSlice()
	// fmt.Println(s)
	// InsertSort(s)
	// fmt.Println(s)

	// s = makeSlice()
	// fmt.Println(s)
	// SelectSort(s)
	// fmt.Println(s)

	s := makeSlice()
	fmt.Println(s)
	s = mergeSort(s)
	if isSorted(s) == false {
		panic("unsorted")
	}
}
