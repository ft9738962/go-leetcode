package main

import "fmt"

// ListNode ...
// type stack struct {
// 	vals []int
// }

func isValidSudoku(board [][]byte) bool {
	hlineCheck := map[int]bool{}
	vlineCheck := map[int]bool{}
	blockCheck := [9]map[int]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if string(board[i][j]) == "." {
				continue
			} else if hlineCheck[int(board[i][j])] == false {
				hlineCheck[int(board[i][j])] = true
			} else {
				return false
			}

			if string(board[j][i]) == "." {
				continue
			} else if vlineCheck[int(board[j][i])] == false {
				vlineCheck[int(board[j][i])] = true
			} else {
				return false
			}

			if blockCheck[i*3+j][int(board[i][j])] == false {
				blockCheck[i*3+j][int(board[i][j])] = true
			} else {
				return false
			}
		}
	}
	return true
}

func main() {

	in := [][]byte{
		[]byte{"8", "3", ".", ".", "7", ".", ".", ".", "."},
	}
	in2 := []int{5, 7, 7, 8, 8, 10}
	in3 := []int{}
	// in4 := ")(((((()())()()))()(()))("
	// in5 := ")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"
	// fmt.Println(longestValidParentheses(in))

	fmt.Println(searchRange(in, 1))
	fmt.Println(searchRange(in2, 8))
	fmt.Println(searchRange(in3, 0))
	// fmt.Println(longestValidParentheses(in4))
	// fmt.Println(longestValidParentheses(in5))

	// fmt.Println(r2)

}
