package main

import (
	mylist "linklist/list"
)

func main() {
	l := mylist.NewList(1, 2, 3, 4, 5)
	mylist.Print(l)
	// l1 := list.New() "container/list" 提供了一个双向链表
	first := mylist.NewList(1, 3, 5)
	second := mylist.NewList(2, 4, 6)
	mylist.Print(first)
	mylist.Print(second)
	third := mylist.MergeList(first, second)
	mylist.Print(third)
}
