package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, cur *ListNode
	if l1 == head {
		return l2
	}
	if l2 == head {
		return l1
	}
	if comp(l1.Val, l2.Val) {
		head = l2
		fmt.Println(head)
	} else {
		head = l1
	}
	cur = head

	for l1.Next != nil && l2.Next != nil {
		if l1.Next == nil {
			cur.Next = l2
			break
		}
		if l2.Next == nil {
			cur.Next = l1
			break
		}
		if comp(l1.Val, l2.Val) {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		}
	}
	return head
}

func comp(a, b int) bool {
	if a > b {
		return true
	}
	return false
}

func main() {
	c := &ListNode{
		Val:  4,
		Next: nil,
	}

	b := &ListNode{
		Val:  2,
		Next: c,
	}

	a := &ListNode{
		Val:  1,
		Next: b,
	}

	f := &ListNode{
		Val:  4,
		Next: nil,
	}

	e := &ListNode{
		Val:  3,
		Next: f,
	}

	d := &ListNode{
		Val:  1,
		Next: e,
	}

	fmt.Println(mergeTwoLists(a, d))

}
