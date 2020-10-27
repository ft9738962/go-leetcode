package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head.Next == nil {
		return head
	}
	ct := countListNodes(head)
	temp := &ListNode{Val: -1, Next: head}
	head = temp
	prev := head
	var left, right, next, first *ListNode
	for i := 0; i < ct/k; i++ {
		left = prev.Next
		right = left.Next
		for j := 0; j < k-1; j++ {
			next = right.Next
			right.Next = left
			left = right
			right = next
		}
		first = prev.Next
		first.Next = right
		prev.Next = left
		prev = first
	}

	return head.Next
}

func countListNodes(head *ListNode) int {
	ct := 0
	for head != nil {
		ct++
		head = head.Next
	}
	return ct
}

func printVals(head *ListNode) {
	fmt.Println("start")
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("end")
}

func main() {
	a1 := ListNode{
		Val:  1,
		Next: nil,
	}

	a2 := ListNode{
		Val:  2,
		Next: nil,
	}

	a3 := ListNode{
		Val:  3,
		Next: nil,
	}

	a4 := ListNode{
		Val:  4,
		Next: nil,
	}

	a5 := ListNode{
		Val:  5,
		Next: nil,
	}
	a6 := ListNode{
		Val:  6,
		Next: nil,
	}

	// var c1 *ListNode

	// d1 := ListNode{
	// 	Val:  2,
	// 	Next: nil,
	// }

	// d2 := ListNode{
	// 	Val:  6,
	// 	Next: nil,
	// }

	// d3 := ListNode{
	// 	Val:  1130,
	// 	Next: nil,
	// }

	a1.Next = &a2
	a2.Next = &a3
	a3.Next = &a4
	a4.Next = &a5
	a5.Next = &a6

	// d1.Next = &d2
	// d2.Next = &d3

	// lists := []*ListNode{&a1, &b1, c1, &d1}
	// lists := []*ListNode{&a1, &b1, &d1}
	// lists := []*ListNode{}

	// var emp *ListNode

	// for _, n := range lists {
	// 	for n.Next != nil {
	// 		fmt.Println(n.Val)
	// 		n = n.Next
	// 	}
	// 	fmt.Println(n.Val)

	// head := mergeTwoLists(&a1, &b1)
	head := reverseKGroup(&a1, 1)
	printVals(head)

}
