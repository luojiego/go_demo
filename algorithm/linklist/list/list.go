package list

import "fmt"

type List struct {
	Node int
	Next *List
}

func NewList(val ...int) *List {
	// 如果为空，则返回空链接
	if len(val) == 0 {
		return nil
	}

	l := &List{
		Node: val[0],
		Next: nil,
	}

	tmp := l
	for i := 1; i < len(val); i++ {
		node := &List{Node: val[i], Next: nil}
		tmp.Next = node
		tmp = node
	}
	return l
}

func Print(l *List) {
	for l != nil {
		fmt.Printf("{val: %d}->", l.Node)
		l = l.Next
	}
	fmt.Println("{nil}")
}
