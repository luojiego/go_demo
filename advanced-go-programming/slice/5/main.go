package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5] //[2,3,4] len=3 cap=8
	s2 := s[2:6:7] //len=4 cap=5

	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))

	s2 = append(s2, 100)
	s2 = append(s2, 100)
	//s2 2,3,4,5,100,100
	s1[2] = 20
	fmt.Println(s1) //[2,3,20]
	fmt.Println(s2) //[2,3,4,5,100,100]
	fmt.Println(s) //[0,1,2,3,20,5,100,7,8,9]
}
