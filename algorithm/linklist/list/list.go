package list

import (
	"fmt"
)

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

func NewCycleList(cycleIndex int, val ...int) *List {
	head := &List{Node: -1, Next: nil}

	node := head
	var cycleNode *List
	for i := 0; i < len(val); i++ {
		node.Next = &List{Node: val[i], Next: nil}
		node = node.Next
		if i == cycleIndex {
			cycleNode = node
		}
	}

	node.Next = cycleNode
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

// 判断链表是否有环
// 使用快慢指针，如果有环，快指针一定会追上慢指针
func HasCycle(head *List) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func HasCycleUseHash(head *List) bool {
	h := make(map[*List]struct{})
	for head != nil {
		if _, ok := h[head]; ok {
			return true
		}
		h[head] = struct{}{}
		head = head.Next
	}
	return false
}

func ReverseLinkList(head *List) *List {
	var pre *List
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 递归实现
func recur(cur, pre *List) *List {
	if cur == nil {
		return pre
	}
	res := recur(cur.Next, cur)
	cur.Next = pre
	return res
}

// 递归实现翻转链表
func ReverseLinkListByRecur(head *List) *List {
	return recur(head, nil)
}

func MiddleNode(head *List) *List {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func KThToLast(head *List, k int) *List {
	pre, cur := head, head
	for i := 0; i < k; i++ {
		if cur.Next == nil {
			return pre
		}
		cur = cur.Next
	}
	for cur != nil {
		pre = pre.Next
		cur = cur.Next
	}
	return pre
}
