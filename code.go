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

	for i := 0; j < len(vals); j++ {
		if i > 0 {
			if order[i] == order[i-1] {
				continue
			}
		}
		nodes := valMap[order[j]]
		for k := 0; k < len(nodes); k++ {
			lists[ct] = nodes[k]
			ct++
		}
	}

	var resultHead, curResult *ListNode

	return resultHead
}

func firstOrderLists(lists []*ListNode, emp *ListNode) ([]int, map[int][]*ListNode, int) {
	vals := []int{}
	valMap := map[int][]*ListNode{}
	nonEmptyListCount := 0
	var list *ListNode
	for i := 0; i < len(lists); i++ {
		if lists[i] != emp {
			list = lists[i]
			for list.Next != nil {
				vals = append(vals, list.Val)
				if _, ok := valMap[list.Val]; ok != false {
					valMap[list.Val] = append(valMap[lists[i].Val], list)
				} else {
					valMap[list.Val] = []*ListNode{list}
				}
				list = list.Next

			}
			vals = append(vals, list.Val)
			if _, ok := valMap[list.Val]; ok != false {
				valMap[list.Val] = append(valMap[lists[i].Val], list)
			} else {
				valMap[list.Val] = []*ListNode{list}
			}
			nonEmptyListCount++
		}
	}
	if nonEmptyListCount == 0 {
		return vals, valMap, nonEmptyListCount
	}
}

func addToMap(valMap map[int][]*ListNode, node *ListNode) map[int][]*ListNode {

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
		Val:  2,
		Next: nil,
	}

	// a2 := ListNode{
	// 	Val:  4,
	// 	Next: nil,
	// }

	// a3 := ListNode{
	// 	Val:  5,
	// 	Next: nil,
	// }

	// b1 := ListNode{
	// 	Val:  1,
	// 	Next: nil,
	// }

	// b2 := ListNode{
	// 	Val:  3,
	// 	Next: nil,
	// }

	// b3 := ListNode{
	// 	Val:  4,
	// 	Next: nil,
	// }

	// b4 := ListNode{
	// 	Val:  700,
	// 	Next: nil,
	// }

	var c1 *ListNode

	d1 := ListNode{
		Val:  -1,
		Next: nil,
	}

	// d2 := ListNode{
	// 	Val:  6,
	// 	Next: nil,
	// }

	// d3 := ListNode{
	// 	Val:  1130,
	// 	Next: nil,
	// }

	// a1.Next = &a2
	// a2.Next = &a3
	// b1.Next = &b2
	// b2.Next = &b3
	// b3.Next = &b4
	// d1.Next = &d2
	// d2.Next = &d3

	// lists := []*ListNode{&a1, &b1, &c1, &d1}
	lists := []*ListNode{&a1, c1, &d1}

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
