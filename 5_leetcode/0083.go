package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	Print(deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{6, nil}}}}))
	//Print(&ListNode{1, &ListNode{1, &ListNode{2, nil}}})
}

func deleteDuplicates(head *ListNode) *ListNode {
	val := head.Val
	var h *ListNode = head

	for h != nil {
		if h.Val != val {
			val = h.Val
		} else if h.Next.Next != nil {
			head.Next = h.Next.Next
		}
		h = h.Next

	}
	return head
}

func Print(l *ListNode) {
	var head *ListNode = l

	fmt.Print("[")

	for head != nil {
		fmt.Print(head.Val, ", ")
		head = head.Next
	}

	fmt.Print("]")
}
