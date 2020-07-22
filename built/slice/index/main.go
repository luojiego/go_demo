package main

import "fmt"

//slice 索引研究

//原因：在遍历删除的时候会出现这一句 s = append(s[:i], s[i+1:]...)
//疑惑：那么当i是最后一个元素的时候，s[i+1:]会不会是一个正常的 slice 呢？或者说会不会panic

func main() {
	//1 nil slice 的访问必然导致会 panic

	//var s []int
	//fmt.Println(s[0])
	//panic: runtime error: index out of range [0] with length 0

	//2 非空的 slice访问呢 len == cap
	//s := []int {1}
	//fmt.Printf("s: %+v, len=%d, cap=%d\n", s, len(s), cap(s)) //s: [1], len=1, cap=1
	////我们来访问第2个位置上的元素 因为只有一个元素 2所在的位置显然是不合法的
	//fmt.Println(s[1]) //panic: runtime error: index out of range [1] with length 1

	//3 非空的 slice 访问 len < cap
	s := make([]int, 2, 4)
	fmt.Printf("s: %+v, len=%d, cap=%d\n", s, len(s), cap(s)) //s: [1], len=1, cap=1
	//我们来访问第3个位置上的元素
	//fmt.Println(s[2]) //panic: runtime error: index out of range [2] with length 2

	s1 := s[2:]
	fmt.Printf("s1: %+v, len=%d, cap=%d\n", s1, len(s1), cap(s1)) //s1: [0 0], len=2, cap=2

	s2 := s[3:]
	fmt.Printf("s2: %+v, len=%d, cap=%d\n", s2, len(s2), cap(s2)) //panic: runtime error: slice bounds out of range [3:2]

}
