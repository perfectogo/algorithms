package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}
	list2 := ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}
	fmt.Println(list1)
	fmt.Println(mergeTwoLists(&list1, &list2).Next.Val)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{}
	dummy.Next = nil
	tail := &dummy

	for l1 != nil || l2 != nil {
		if l1 == nil {
			(*tail).Next = l2
			break
		} else if l2 == nil {
			(*tail).Next = l1
			break
		} else {
			if l1.Val <= l2.Val {
				(*tail).Next = l1
				l1 = (*l1).Next
			} else {
				(*tail).Next = l2
				l2 = (*l2).Next
			}
			tail = tail.Next
		}
	}

	return dummy.Next
}
