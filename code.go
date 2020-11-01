package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func divide(dividend int, divisor int) int {
	var sign, result int
	sign = 1

	if dividend == 0 {
		return 0
	} else if divisor == 1 {
		result = dividend
	} else if divisor == -1 {
		result = -1 * dividend
	}
	if dividend > 0 && divisor < 0 {
		sign = -1
		for dividend > 0 {

		}
	} else if dividend >= 0 && divisor < 0 {

	} else {

<<<<<<< HEAD
	}

	if result > 1<<31-1 || result < -(1<<31) {
		return result * sign
	}
=======
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
>>>>>>> 79e80527b2aefb6271d4665f37dfff5355375d81
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
<<<<<<< HEAD
	test1 := []int{1, 1}
	test2 := []int{1, 1, 2}
	test3 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// a6 := ListNode{
	// 	Val:  6,
	// 	Next: nil,
	// }
=======
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
>>>>>>> 79e80527b2aefb6271d4665f37dfff5355375d81

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

<<<<<<< HEAD
	// a5.Next = &a6
=======
	a1.Next = &a2
	a2.Next = &a3
	a3.Next = &a4
	a4.Next = &a5
	a5.Next = &a6
>>>>>>> 79e80527b2aefb6271d4665f37dfff5355375d81

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
<<<<<<< HEAD
	answer1 := test1[:removeDuplicates(test1)]
	answer2 := test2[:removeDuplicates(test2)]
	answer3 := test3[:removeDuplicates(test3)]
	fmt.Printf("answer1 %v\n", answer1)
	fmt.Printf("answer2 %v\n", answer2)
	fmt.Printf("answer3 %v\n", answer3)
=======
	head := reverseKGroup(&a1, 1)
	printVals(head)
>>>>>>> 79e80527b2aefb6271d4665f37dfff5355375d81

}
