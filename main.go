package main

import (
	"fmt"
	"sort"
)

// ListNode ...
// type stack struct {
// 	vals []int
// }

func combinationSum(candidates []int, target int) [][]int {
	var solutions [][]int
	var solution []int
	sort.Ints(candidates)
	fmt.Println(candidates)
	recurCombine(candidates, target, solution,
		&solutions)
	return solutions
}

func recurCombine(candidates []int, target int, solution []int,
	solutionsPtr *[][]int) {
	for i := len(candidates) - 1; i > -1; i-- {
		v := candidates[i]
		newTarget := target - v
		if newTarget == 0 {
			// fmt.Println(candidates[i], solution, newTarget, append(solution, candidates[i]))
			solutionCopy := make([]int, len(solution)+1)
			copy(solutionCopy, solution)
			solutionCopy[len(solutionCopy)-1] = v
			*solutionsPtr = append(*solutionsPtr, solutionCopy)
		} else if newTarget > 0 {
			recurCombine(candidates[:i+1], newTarget, append(solution, v),
				solutionsPtr)
		} else {
		}
	}
}

func main() {

	in := []int{2, 7, 6, 3, 5, 1}
	// in2 := []int{5, 7, 7, 8, 8, 10}
	// in3 := []int{}
	// in4 := ")(((((()())()()))()(()))("
	// in5 := ")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"
	// fmt.Println(longestValidParentheses(in))

	fmt.Println(combinationSum(in, 9))
	// combinationSum(in, 9)
	// fmt.Println(searchRange(in2, 8))
	// fmt.Println(searchRange(in3, 0))
	// fmt.Println(longestValidParentheses(in4))
	// fmt.Println(longestValidParentheses(in5))

	// fmt.Println(r2)

}
