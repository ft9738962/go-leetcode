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
	k := len(lists)
	if k == 0 {
		return emp
	} else if k == 1 {
		return lists[0]
	}

	order := []int{}
	firstMap := map[int][]*ListNode{}
	for i := 0; i < k; i++ {
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
	fmt.Println(len(order))
	ct := 0
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
			fmt.Println("ct:", ct)
		}
	}
	lists = lists[:ct]

	var resultHead, curResult *ListNode
	resultHead = lists[0]
	curResult = resultHead
	for len(lists) > 1 {
		if lists[0].Next == nil {
			lists = lists[1:]
			curResult.Next = lists[0]
		}
		for lists[0].Next != nil {
			if lists[0].Next.Val <= lists[1].Val {
				curResult = curResult.Next
				lists[0] = lists[0].Next
				continue
			}
			break
		}
		lists = reorderLists(lists)
	}
	curResult.Next = lists[0]
	return resultHead
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
			_ = copy(lists[1:mid], lists[0:mid-1])
			lists[mid-1] = n
			return lists
		}
	}
	if lists[right].Val < n.Val {
		_ = copy(lists[1:right+1], lists[:right])
		lists[right] = n
	} else {
		_ = copy(lists[1:right], lists[:right-1])
		lists[right-1] = n
	}
	return lists
}

func main() {
	// result := generateParenthesis(4)
	// fmt.Println(result)

}
