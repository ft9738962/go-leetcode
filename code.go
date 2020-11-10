package main

import "fmt"

// ListNode ...
type stack struct {
	vals []int
}

func search(nums []int, target int) int {
	l, m, r := 0, len(nums)/2, len(nums)-1
	for l != m || m != r {
		if target > nums[l] && target < nums[m] {
			return oneDirectionSearch(nums, target, l, m)
		} else if target > nums[l] && target > nums[m] {
			l, m, r = m, (m+r)/2, r
		} else if target < nums[l] && target > nums[m] {
			l, m, r = m, (m+r)/2, r
		} else if target < nums[l] && target < nums[m] {
			return oneDirectionSearch(nums, target, l, m)
		} else if target == nums[l] {
			return l
		} else if target == nums[m] {
			return m
		} else if target == nums[r] {
			return r
		}
	}
	return -1
}

func oneDirectionSearch(nums []int, target int, l, r int) int {

}

func main() {

	// in := ")()())()()()(())))()()()()"
	// in2 := "()(())"
	in3 := ")()())"
	in4 := ")(((((()())()()))()(()))("
	in5 := ")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"
	// fmt.Println(longestValidParentheses(in))

	// fmt.Println(longestValidParentheses(in2))
	fmt.Println(longestValidParentheses(in3))
	fmt.Println(longestValidParentheses(in4))
	fmt.Println(longestValidParentheses(in5))

	// fmt.Println(r2)

}
