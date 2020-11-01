package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicates(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	nums = nums[:j]
	return j + 1
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
	test1 := []int{1, 1}
	test2 := []int{1, 1, 2}
	test3 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// a6 := ListNode{
	// 	Val:  6,
	// 	Next: nil,
	// }

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

	// a5.Next = &a6

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
	answer1 := test1[:removeDuplicates(test1)]
	answer2 := test2[:removeDuplicates(test2)]
	answer3 := test3[:removeDuplicates(test3)]
	fmt.Printf("answer1 %v\n", answer1)
	fmt.Printf("answer2 %v\n", answer2)
	fmt.Printf("answer3 %v\n", answer3)

}
