package main

import "fmt"

//空切片和nil切片的区分
func main() {
	var s1 []int		//nil slice
	var s2 = []int{}	//empty slice

	fmt.Printf("s1:%#v cap(s1)=%d len(s1)=%d\n", s1, cap(s1), len(s1))
	fmt.Printf("s2:%#v cap(s2)=%d len(s2)=%d\n", s2, cap(s2), len(s2))


}
