package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 首先难点在于怎么开始，怎么结束
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := l1
	q := l2
	cur := result
	carry := 0
	for p != nil || q != nil {
		x := 0
		if p != nil {
			x = p.Val
		}

		y := 0
		if q != nil {
			y = q.Val
		}

		sum := x + y + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		if p != nil {
			p = p.Next
		}

		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return result.Next
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x > 0 && x < 10 {
		return true
	}

	// 先将 x 转换为 slice
	var s []int
	for {
		s = append(s, x%10)
		x /= 10
		// 如果 x < 小于 10 时，应该退出循环
		if x < 10 {
			s = append(s, x)
			break
		}
	}
	// fmt.Println(s)

	i := 0          // 从 slice 第一个位置
	j := len(s) - 1 // slice 最后一个位置
	for {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
		if i == j || i > j {
			break
		}
	}
	return true
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := l1
	q := l2
	cur := result

	for p != nil && q != nil {
		x := p.Val
		y := q.Val

		if x > y {
			cur.Val = y
			q = q.Next
			cur.Next = &ListNode{}
		} else if x < y {
			cur.Val = x
			p = p.Next
			cur.Next = &ListNode{}
		} else {
			cur.Val = x
			p = p.Next
			q = q.Next
			cur.Next = &ListNode{Val: y}
			if p == nil && q == nil {
				return result
			} else {
				cur.Next.Next = &ListNode{}
			}
			cur = cur.Next

		}
		cur = cur.Next
	}

	if p != nil {
		for p != nil {
			cur.Val = p.Val
			if p.Next != nil {
				cur.Next = &ListNode{}
				cur = cur.Next
			}
			p = p.Next
		}
	}

	if q != nil {
		for q != nil {
			cur.Val = q.Val
			if q.Next != nil {
				cur.Next = &ListNode{}
				cur = cur.Next
			}
			q = q.Next
		}
	}
	return result
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(1))

	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	result := mergeTwoLists(l1, l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
