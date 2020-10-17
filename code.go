package main

import (
	"fmt"
)

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	dist := 1
	var before = head
	for cur.Next != nil {
		if dist >= n+1 {
			before = before.Next
		}
		cur = cur.Next
		dist++
	}
	if dist == n {
		head = before.Next
	} else {
		before.Next = before.Next.Next
	}
	return head
}

func main() {
	test1 := []int{1, 0, -1, 0, -2, 2}
	test2 := []int{-1, -5, -5, -3, 2, 5, 0, 4}
	test3 := []int{-1, 2, 2, -5, 0, -1, 4}
	test4 := []int{2, -4, -5, -2, -3, -5, 0, 4, -2}
	test5 := []int{-2, -1, -1, 1, 1, 2, 2}
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 100; i++ {
	// 	test1 = append(test1, rand.Intn(101)-50)
	// }
	fmt.Println(fourSum(test1, 0))
	fmt.Println(fourSum(test2, -7))
	fmt.Println(fourSum(test3, 3))
	fmt.Println(fourSum(test4, -14))
	fmt.Println(fourSum(test5, 0))
}
