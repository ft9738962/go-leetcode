package main

import (
	"fmt"
	"sort"
)

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	var emp *ListNode
	if len(lists) == 0 {
		return emp
	}

	vals, valMap, nonEmptyListCount := firstOrderLists(lists, emp)
	if nonEmptyListCount == 0 {
		return emp
	}

	sort.Ints(vals)
	var resultHead, curResult *ListNode

	for i := 0; i < len(vals); i++ {
		if i > 0 {
			if vals[i] == vals[i-1] {
				continue
			}
		}

		nodes := valMap[vals[i]]
		for j := 0; j < len(nodes); j++ {
			if resultHead == emp {
				resultHead = nodes[j]
				curResult = resultHead
			} else {
				curResult.Next = nodes[j]
				curResult = curResult.Next
			}
		}
	}

	return resultHead
}

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	prehead := &ListNode{Val: 0}
	cur := prehead

	for (a.Next != nil && a.Val >= b.Val) || 
		(b.Next != nil && a.Val <= b.Val) {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}

	if a.Next == nil {
			cur.Next = b
			b = b.Next
			cur = cur.Next
		}
		cur.Next = a
		cur.Next = 
	}
}

func printVals(head *ListNode) {
	fmt.Println("start")
	for head.Next != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println(head.Val)
	fmt.Println("end")
}

func main() {
	a1 := ListNode{
		Val:  1,
		Next: nil,
	}

	a2 := ListNode{
		Val:  4,
		Next: nil,
	}

	a3 := ListNode{
		Val:  5,
		Next: nil,
	}

	b1 := ListNode{
		Val:  1,
		Next: nil,
	}

	b2 := ListNode{
		Val:  3,
		Next: nil,
	}

	b3 := ListNode{
		Val:  4,
		Next: nil,
	}

	// b4 := ListNode{
	// 	Val:  700,
	// 	Next: nil,
	// }

	// var c1 *ListNode

	d1 := ListNode{
		Val:  2,
		Next: nil,
	}

	d2 := ListNode{
		Val:  6,
		Next: nil,
	}

	// d3 := ListNode{
	// 	Val:  1130,
	// 	Next: nil,
	// }

	a1.Next = &a2
	a2.Next = &a3
	b1.Next = &b2
	b2.Next = &b3
	// b3.Next = &b4
	d1.Next = &d2
	// d2.Next = &d3

	// lists := []*ListNode{&a1, &b1, c1, &d1}
	lists := []*ListNode{&a1, &b1, &d1}

	// var emp *ListNode

	// for _, n := range lists {
	// 	for n.Next != nil {
	// 		fmt.Println(n.Val)
	// 		n = n.Next
	// 	}
	// 	fmt.Println(n.Val)

	head := mergeKLists(lists)
	printVals(head)

}
