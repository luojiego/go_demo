package main

import "fmt"

func main() {
	//len:5
	//cap:10
	s := make([]int, 5, 10)
	s[2] = 5
	fmt.Println(s)
	fmt.Println("len(s)=",len(s),"\tcap(s)=",cap(s))
}
