package main

import (
	"fmt"
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

	hasCycleList := mylist.NewCycleList(2, 6, 7, 8, 9, 10)
	fmt.Printf("has cycle: %t\n", mylist.HasCycle(hasCycleList))
	fmt.Printf("has cycle: %t\n", mylist.HasCycle(first))
	fmt.Printf("has cycle use hash: %t\n", mylist.HasCycleUseHash(hasCycleList))
	fmt.Printf("has cycle use hash: %t\n", mylist.HasCycleUseHash(first))

	mylist.ReverseLinkList(mylist.NewList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)).Print()
	mylist.ReverseLinkListByRecur(mylist.NewList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)).Print()

	mylist.MiddleNode(mylist.NewList(1, 2, 3, 4, 5)).Print()
	mylist.MiddleNode(mylist.NewList(1, 2, 3, 4, 5, 6)).Print()

	mylist.KThToLast(mylist.NewList(1, 2, 3, 4, 5), 2).Print()
	mylist.KThToLast(mylist.NewList(1, 2, 3, 4, 5, 6), 5).Print()
}
