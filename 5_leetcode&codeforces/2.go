package main

import "fmt"

func main() {
	l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 := ListNode{9, &ListNode{9, &ListNode{9, nil}}}

	l3 := addTwoNumbers(&l1, &l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		dummyHead   = &ListNode{}
		current     = dummyHead
		x, y, carry int
	)

	for l1 != nil || l2 != nil || carry != 0 {

		if l1 != nil {
			x = l1.Val
		} else {
			x = 0
		}
		//fmt.Println(x)
		if l2 != nil {
			y = l2.Val
		} else {
			y = 0
		}
		//fmt.Println(y)
		sum := carry + x + y
		carry = sum / 10
		current.Next = &ListNode{sum % 10, nil}
		current = current.Next

		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}

	}
	return dummyHead.Next
}
