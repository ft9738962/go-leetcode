package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	a := head
	b := head.Next
	head = b
	var last *ListNode
	last = nil

	for {
		if last == nil {
			last = a
		} else {
			last.Next = b
		}
		a.Next, b.Next = b.Next, a
		if a.Next == nil {
			break
		}
		if a.Next.Next == nil {
			break
		}
		last, a, b = a, a.Next, a.Next.Next
	}

	return head
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
	head := swapPairs(&a1)
	printVals(head)

}
