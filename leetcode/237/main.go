package main

import "fmt"

/**
 * Definition for singly-linked list.
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	// p := node.Next.Next
	// node = node.Next.Next
	*node = *node.Next
	fmt.Printf("%d:{%p}\n", node.Val, node)
	fmt.Printf("%d:{%p}\n", node.Next.Val, node.Next)
	// fmt.Printf("%d:{%p}\n", node.Next.Next.Val, node.Next.Next)
}

func print(node *ListNode) {
	p := node
	for p != nil {
		fmt.Printf("%d:{%p}", p.Val, p)
		if p.Next != nil {
			fmt.Printf("->")
		} else {
			break
		}
		p = p.Next
	}
	fmt.Println()
}

func find(n int, node *ListNode) *ListNode {
	p := node
	for p != nil {
		if p.Val == n {
			return p
		}
		p = p.Next
	}
	return nil
}

func main() {
	l := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	print(l)

	deleteNode(find(5, l))

	print(l)
}
