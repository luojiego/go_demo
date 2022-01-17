package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5] //s1[?] len(s1)=? cap(s1)=?
	s2 := s1[2:6:7] //s2[?] len(s2)=? cap(s2)=?

	s2 = append(s2, 100) //s2[?] len(s2)=? cap(s2)=?
	s2 = append(s2, 200) //s2[?] len(s2)=? cap(s2)=?

	s1[2] = 20//s1[?] len(s1)=? cap(s1)=?

	fmt.Println(s1) //[?]
	fmt.Println(s2) //[?]
	fmt.Println(s) //[?]
}
