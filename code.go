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
	} else if len(lists) == 1 {
		return lists[0]
	}

	lists = firstOrderLists(lists, emp)

	var resultHead, curResult *ListNode
	resultHead = lists[0]
	curResult = resultHead
	for len(lists) > 1 {

		if lists[0].Next == nil {
			lists = lists[1:]
			continue
		}
		for lists[0].Next != nil {
			lists[0] = lists[0].Next
			if lists[0].Val <= lists[1].Val {
				curResult = curResult.Next
				continue
			}
			break
		}
		lists = reorderLists(lists)
		curResult.Next = lists[0]
		curResult = curResult.Next
	}
	return resultHead
}

func firstOrderLists(lists []*ListNode, emp *ListNode) []*ListNode {
	ct := 0
	order := []int{}
	firstMap := map[int][]*ListNode{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != emp {
			order = append(order, lists[i].Val)
			if _, ok := firstMap[lists[i].Val]; ok != false {
				firstMap[lists[i].Val] = append(firstMap[lists[i].Val], lists[i])
			} else {
				firstMap[lists[i].Val] = []*ListNode{lists[i]}
			}
		}
	}
	sort.Ints(order)

	for j := 0; j < len(order); j++ {
		if j > 0 {
			if order[j] == order[j-1] {
				continue
			}
		}
		nodes := firstMap[order[j]]
		for k := 0; k < len(nodes); k++ {
			lists[ct] = nodes[k]
			ct++
		}
	}
	return lists[:ct]
}

func reorderLists(lists []*ListNode) []*ListNode {
	n := lists[0]
	left, right, mid := 1, len(lists)-1, 0
	for left < right {
		mid = (left + right) / 2
		if n.Val > lists[mid].Val {
			left = mid + 1
		} else if n.Val < lists[mid].Val {
			right = mid
		} else {
			_ = copy(lists[0:mid-1], lists[1:mid])
			lists[mid-1] = n
			return lists
		}
	}
	if lists[right].Val < n.Val {
		_ = copy(lists[:right], lists[1:right+1])
		lists[right] = n
	} else {
		_ = copy(lists[:right-1], lists[1:right])
		lists[right-1] = n
	}
	return lists
}

func main() {
	a1 := ListNode{
		Val:  100,
		Next: nil,
	}

	a2 := ListNode{
		Val:  800,
		Next: nil,
	}

	a3 := ListNode{
		Val:  1000,
		Next: nil,
	}

	b1 := ListNode{
		Val:  70,
		Next: nil,
	}

	b2 := ListNode{
		Val:  90,
		Next: nil,
	}

	b3 := ListNode{
		Val:  130,
		Next: nil,
	}

	b4 := ListNode{
		Val:  700,
		Next: nil,
	}

	c1 := ListNode{
		Val:  110,
		Next: nil,
	}

	d1 := ListNode{
		Val:  130,
		Next: nil,
	}

	d2 := ListNode{
		Val:  830,
		Next: nil,
	}

	d3 := ListNode{
		Val:  1130,
		Next: nil,
	}

	a1.Next = &a2
	a2.Next = &a3
	b1.Next = &b2
	b2.Next = &b3
	b3.Next = &b4
	d1.Next = &d2
	d2.Next = &d3

	lists := []*ListNode{&a1, &b1, &c1, &d1}

	// var emp *ListNode

	// for _, n := range lists {
	// 	for n.Next != nil {
	// 		fmt.Println(n.Val)
	// 		n = n.Next
	// 	}
	// 	fmt.Println(n.Val)

	head := mergeKLists(lists)
	a := 0
	fmt.Println("start")
	for head.Next != nil && a < 15 {
		fmt.Println(head.Val)
		head = head.Next
		a++
	}
	fmt.Println(head.Val)
	fmt.Println("end")
}
