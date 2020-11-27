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

	return solutions
}

func try(candidates []int, target int, solution []int, solutions *[][]int, recMap *map[int]bool) {
	for i := len(candidates) - 1; i >= 0; i-- {
		newTarget := target - candidates[i]
		if ok := recMap[newTarget] {
			
		}
		if newTarget == 0 {
			solution = append(solution, candidates[i])
			*solutions = append(*solutions, solution)
			break
		} else if newTarget > 0 {
			&recMap[newTarget] = 1
			try(candidates, newTarget, append(solution, candidates[i]), solutions)
		} else {
			continue
		}
	}
}

func main() {

	in := [][]byte{
		{53, 51, 46, 46, 55, 46, 46, 46, 46},
		{54, 46, 46, 49, 57, 53, 46, 46, 46},
		{46, 57, 56, 46, 46, 46, 46, 54, 46},
		{56, 46, 46, 46, 54, 46, 46, 46, 51},
		{52, 46, 46, 56, 46, 51, 46, 46, 49},
		{55, 46, 46, 46, 50, 46, 46, 46, 54},
		{46, 54, 46, 46, 46, 46, 50, 56, 46},
		{46, 46, 46, 52, 49, 57, 46, 46, 53},
		{46, 46, 46, 46, 56, 46, 46, 55, 57},
	}
	// in2 := []int{5, 7, 7, 8, 8, 10}
	// in3 := []int{}
	// in4 := ")(((((()())()()))()(()))("
	// in5 := ")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"
	// fmt.Println(longestValidParentheses(in))

	fmt.Println(isValidSudoku(in))
	// fmt.Println(searchRange(in2, 8))
	// fmt.Println(searchRange(in3, 0))
	// fmt.Println(longestValidParentheses(in4))
	// fmt.Println(longestValidParentheses(in5))

	// fmt.Println(r2)

}
