package main

import "fmt"

func main() {
	//使用8:100 表示第8个元素为100 中间的则被初始为0
	s1 := []int{0,1,2,3,8:100}
	fmt.Println(s1, len(s1), cap(s1))

	//换成100:8 则说明第100个元素为8 cap和len均为100
	s2 := []int{0,1,2,3,100:8}
	fmt.Println(s2, len(s2), cap(s2))

}
