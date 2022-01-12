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
	var c1 int
	for i := 0; i < sliceSize; i++ {
		s[i] = rand.Intn(100)
		c1 += s[i]
	}
	// fmt.Println(s)
	// 执行排序
	s = quickSort(s)
	if isSorted(s) {
		return s
	}
	var c2 int
	for _, v := range s {
		c2 += v
	}
	if c1 != c2 {
		panic("error occurs")
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

// pretty print
func printSlice(s []int) {
	fmt.Printf("   ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%2d ", i)
	}
	for i := 0; i < len(s); i++ {
		if i%10 == 0 {
			fmt.Printf("\n%d: ", i/10)
		}
		fmt.Printf("%2d ", s[i])
	}
	fmt.Println()
}

// 	   0  1  2  3  4  5  6  7  8  9
// 0:  0  0  2  2  2  3  3  5  5  5
// 1:  6  7  8 10 11 11 13 15 18 18
// 2: 20 21 23 24 25 26 28 28 28 29
// 3: 31 31 33 33 33 36 37 37 37 38
// 4: 40 40 41 41 43 43 45 46 46 47
// 5: 47 47 47 47 51 52 53 53 55 56
// 6: 56 56 57 58 59 59 59 61 62 63
// 7: 63 63 66 66 74 76 77 78 78 81
// 8: 81 83 85 87 87 87 88 88 89 89
// 9: 90 90 91 94 94 94 95 96 98 99

// find first val in s
func findFirstVal(s []int, val int) int {
	low, high := 0, len(s)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if s[mid] == val {
			i := mid
			for ; i > 0; i-- {
				if s[i-1] < val {
					return i
				}
			}
			return i
		} else if s[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// find last val in s
func findLastVal(s []int, val int) int {
	low, high := 0, len(s)-1
	h := high
	for low <= high {
		mid := low + ((high - low) >> 1)
		if s[mid] == val {
			i := mid
			for ; i < h-1; i++ {
				if s[i+1] > val {
					return i
				}
			}
			return i
		} else if s[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// find last gt than val in s
func findFisrtGtVal(s []int, val int) int {
	low, high := 0, len(s)-1
	h := high
	for low <= high {
		mid := low + ((high - low) >> 1)
		if s[mid] == val {
			for ; mid < h-1; mid++ {
				if s[mid+1] > val {
					return mid
				}
			}
			return mid
		} else if s[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// find first lt than val in s
func findLastLtVal(s []int, val int) int {
	low, high := 0, len(s)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if s[mid] == val {
			for ; mid > 0; mid-- {
				if s[mid-1] < val {
					return mid - 1
				}
			}
			return mid
		} else if s[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	s := generateSlice()
	// fmt.Println(s)
	printSlice(s)
	fmt.Println(findFirstVal(s, 41))
	fmt.Println(findLastVal(s, 41))
	fmt.Println(findFisrtGtVal(s, 99))
	fmt.Println(findLastLtVal(s, 0))
	// fmt.Println(bsearch(s, 9947))
	// fmt.Println(recursionBesarch(s, 13))
}
