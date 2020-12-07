package main

import (
	"fmt"
)

// ListNode ...
// type stack struct {
// 	vals []int
// }

func combinationSum(candidates []int, target int) [][]int {
	solutions := [][]int{}
	recMap := make(map[int]bool)
	recurTry(candidates, target, []int{}, &solutions, &recMap)
	return solutions
}

func recurTry(candidates []int, target int, solution []int, solutions *[][]int, recMap *map[int]bool) {
	for i := len(candidates) - 1; i >= 0; i-- {
		newTarget := target - candidates[i]

		if newTarget == 0 {
			solution = append(solution, candidates[i])
			*solutions = append(*solutions, solution)
			solution = []int{}
			continue
		} else if newTarget > 0 {
			if ok := (*recMap)[newTarget]; ok == true {
				continue
			} else {
				solution = append(solution, candidates[i])
				recurTry(candidates[:i+1], newTarget, solution, solutions, recMap)
				(*recMap)[newTarget] = true
			}
		} else {
			if i == 0 {
				solution = []int{}
			}
			break
		}
	}
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
