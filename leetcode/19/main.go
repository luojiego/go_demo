package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	h := head
	head = head.Next
	// 第一次 head 肯定不为 nil
	for head != nil && head != h {
		for p := h; p != head && p != nil; {
			if p.Next == head {
				return true
			}
			p = p.Next
		}
		head = head.Next
	}
	return false
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	ret, remove, last := head, head, head
	for i := 0; i < n; i++ {
		last = last.Next // 2 3
	}

	if last == nil {
		// ret.Next = nil
		return ret.Next
	}

	for last.Next != nil {
		remove = remove.Next // 2 3
		last = last.Next     // 4 5
	}
	// 删除掉 remove->next
	remove.Next = remove.Next.Next
	return ret
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Printf("{val: %d}->", list.Val)
		list = list.Next
	}
	fmt.Println("nil")
}

func main() {
	// list := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	// printList(removeNthFromEnd(list, 2))
	// list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{5, nil}}}}}
	// printList(removeNthFromEnd(list1, 2))
	// list2 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	// printList(removeNthFromEnd(list2, 1))
	// list3 := &ListNode{Val: 1, Next: nil}
	// printList(removeNthFromEnd(list3, 1))
	node := &ListNode{Val: 1}
	node.Next = node
	hasCycle(node)
}
