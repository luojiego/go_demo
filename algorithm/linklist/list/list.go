package list

import "fmt"

type List struct {
	Node int
	Next *List
}

func (l *List) Print() {
	index := 0
	for l != nil {
		fmt.Printf("{val: %d}->", l.Node)
		l = l.Next
		if index++; index > 10 {
			fmt.Println("...")
			break
		}
	}
	fmt.Println("{nil}")
}

func NewList(val ...int) *List {
	head := &List{Node: -1, Next: nil}

	node := head
	for i := 0; i < len(val); i++ {
		node.Next = &List{Node: val[i], Next: nil}
		node = node.Next
	}

	return head.Next
}

// 非递归实现
func MergeList(first, second *List) *List {
	// 创建一个新的链表
	head := NewList(-1)
	node := head
	for first != nil && second != nil {
		if first.Node < second.Node {
			node.Next = first
			first = first.Next
		} else {
			node.Next = second
			second = second.Next
		}
		node = node.Next
	}

	if first != nil {
		node.Next = first
	} else {
		node.Next = second
	}
	return head.Next
}

// 递归实现
func MergeTwoList(first, second *List) *List {
	if first == nil {
		return second
	}

	if second == nil {
		return first
	}

	if first.Node < second.Node {
		first.Next = MergeTwoList(first.Next, second)
		return first
	} else {
		second.Next = MergeTwoList(first, second.Next)
		return second
	}
}
