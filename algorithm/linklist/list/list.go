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

func MergeList(first, second *List) *List {
	p, q := first, second
	// ret := first
	var ret *List
	var tmp *List
	for p != nil && q != nil {
		if ret == nil {
			ret = &List{}
			tmp = ret
		}

		if p.Node < q.Node {
			tmp.Node = p.Node
			tmp.Next = &List{}
			tmp = tmp.Next
			p = p.Next
		} else {
			tmp.Node = q.Node
			tmp.Next = &List{}
			tmp = tmp.Next
			q = q.Next
		}
	}

	for p != nil {
		tmp.Node = p.Node
		if p.Next != nil {
			tmp.Next = &List{}
			tmp = tmp.Next
			p = p.Next
		} else {
			break
		}
	}

	for q != nil {
		tmp.Node = q.Node
		if q.Next != nil {
			tmp.Next = &List{}
			tmp = tmp.Next
			q = q.Next
		} else {
			break
		}
	}
	return ret
}
