package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func longestValidParentheses(s string) int {
	var longest []int
	longest = []int{0}
	var tmp int
	for i := 1; i < len(s); i++ {
		if string(s[i]) == "(" {
			longest = append(longest, 0)
		} else {
			if i-1 >= 0 && string(s[i-1]) == "(" {
				if i == 1 {
					longest = append(longest, 2)
				} else {
					longest = append(longest, longest[i-2]+2)
				}
			} else if i-1 >= 0 && string(s[i-1]) == ")" && i-longest[i-1]-1 >= 0 && string(s[i-longest[i-1]-1]) == "(" {
				if i-longest[i-1]-2 >= 0 {
					tmp = longest[i-longest[i-1]-2]
				} else {
					tmp = 0
				}
				longest = append(longest, longest[i-1]+2+tmp)
			} else {
				longest = append(longest, 0)
			}
		}
	}
	return maxLongest(longest)
}

func maxLongest(longest []int) int {
	maxLen := 0
	for i := 0; i < len(longest); i++ {
		if longest[i] > maxLen {
			maxLen = longest[i]
		}
	}
	return maxLen
}

// 	left := []int32{}
// 	for _, r := range s {
// 		if r == 40 {
// 			left = append(left, r)
// 		} else if len(left) >= 1 {
// 			if r == 41 && left[len(left)-1] == 40 {
// 				left = left[:len(left)-1]
// 			} else {
// 				return false
// 			}
// 		} else {
// 			return false
// 		}
// 	}
// 	if len(left) > 0 {
// 		return false
// 	}
// 	return true
// }

// func printVals(head *ListNode) {
// 	fmt.Println("start")
// 	for head != nil {
// 		fmt.Println(head.Val)
// 		head = head.Next
// 	}
// 	fmt.Println("end")
// }

func main() {

	// in := ")()())()()()(())))()()()()"
	// in2 := "()(())"
	in3 := "(()"
	in4 := ")(((((()())()()))()(()))("
	in5 := ")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"
	// fmt.Println(longestValidParentheses(in))

	// fmt.Println(longestValidParentheses(in2))
	fmt.Println(longestValidParentheses(in3))
	fmt.Println(longestValidParentheses(in4))
	fmt.Println(longestValidParentheses(in5))

	// fmt.Println(r2)

}
