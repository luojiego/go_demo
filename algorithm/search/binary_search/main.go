package main

import (
	"fmt"
	"math/rand"
)

const (
	sliceSize = 100
)

func generateSlice() []int {
	s := make([]int, sliceSize, sliceSize)
	for i := 0; i < sliceSize; i++ {
		s[i] = rand.Intn(10000)
	}
	// 执行排序
	s = quickSort(s)
	if isSorted(s) {
		return s
	}
	panic("error")
}

func quickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	q := partition(s)
	quickSort(s[:q])
	quickSort(s[q+1:])
	return s
}

func partition(s []int) int {
	left, right := 0, len(s)-1
	pivot := rand.Intn(len(s))
	// 使用哨兵来简化判断
	s[pivot], s[right] = s[right], s[pivot]
	for i := range s {
		if s[i] < s[right] {
			s[left], s[i] = s[i], s[left]
			left++
		}
	}
	s[left], s[right] = s[right], s[left]
	return left
}

func isSorted(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

// 递归实现
func recursionBesarch(s []int, value int) int {
	low, high := 0, len(s)-1
	return recursionBesarchBranch(s, low, high, value)
}

func recursionBesarchBranch(s []int, low, high, value int) int {
	if low > high {
		return -1
	}

	mid := low + ((high - low) >> 1)
	if s[mid] == value {
		return mid
	} else if s[mid] < value {
		return recursionBesarchBranch(s, mid+1, high, value)
	} else {
		return recursionBesarchBranch(s, low, mid-1, value)
	}
}

// 非递归实现
func bsearch(s []int, value int) int {
	low, high := 0, len(s)-1

	for low <= high {
		mid := (low + high) / 2
		if s[mid] == value {
			return mid
		} else if s[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	s := generateSlice()
	fmt.Println(len(s))
	fmt.Println(s)
	fmt.Println(bsearch(s, 9947))
	fmt.Println(recursionBesarch(s, 13))
}
