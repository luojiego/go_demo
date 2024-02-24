package main

import (
	mylist "linklist/list"
)

func main() {
	first := mylist.NewList(1, 3, 5, 6)
	second := mylist.NewList(2, 4, 6)

	first.Print()
	second.Print()

	mylist.MergeList(first, second).Print()

	first = mylist.NewList(1, 3, 5, 6)
	second = mylist.NewList(2, 4, 6)

	mylist.MergeTwoList(first, second).Print()
}
