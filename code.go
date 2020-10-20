package main

import (
	"fmt"
)

// ListNode ...
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func generateParenthesis(n int) []string {

	result := []string{}
	recursiveGenerate(&result, "", n, 0, 0, 0)
	return result
}

func recursiveGenerate(result *[]string, current string,
	length, curLength, left, right int) {
	if length == left && length == right {
		*result = append(*result, current)
	} else if left < length && left > right {
		recursiveGenerate(result, current+"(",
			length, curLength+1, left+1, right)
		recursiveGenerate(result, current+")",
			length, curLength+1, left, right+1)
	} else if left == length {
		recursiveGenerate(result, current+")",
			length, curLength+1, left, right+1)
	} else {
		recursiveGenerate(result, current+"(",
			length, curLength+1, left+1, right)
	}
}

func main() {
	result := generateParenthesis(4)
	fmt.Println(result)

}
