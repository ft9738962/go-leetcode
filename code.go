package main

import (
	"fmt"
)

// ListNode ...
// type stack struct {
// 	vals []int
// }

func combinationSum(candidates []int, target int) [][]int {
	solutions := recurCombine(candidates, target)
	return solutions
}

func recurCombine(candidates []int, target int) (solutions [][]int) {
	for i := len(candidates) - 1; i >= 0; i-- {
		newTarget := target - candidates[i]

		if newTarget == 0 {
			answer := []int{candidates[i]}
			return [][]int{answer}
		} else if newTarget > 0 {
			answers := recurCombine(candidates[:i+1], newTarget)
			for _, answer := range answers {
				solutions = append(solutions,
					append(answer, candidates[i]))
			}
			return solutions
		} else {
			break
		}
	}
	return
}

func main() {

	in := []int{2, 3, 6, 7}
	// in2 := []int{5, 7, 7, 8, 8, 10}
	// in3 := []int{}
	// in4 := ")(((((()())()()))()(()))("
	// in5 := ")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"
	// fmt.Println(longestValidParentheses(in))

	fmt.Println(combinationSum(in, 7))
	// fmt.Println(searchRange(in2, 8))
	// fmt.Println(searchRange(in3, 0))
	// fmt.Println(longestValidParentheses(in4))
	// fmt.Println(longestValidParentheses(in5))

	// fmt.Println(r2)

}
