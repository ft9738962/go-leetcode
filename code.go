package main

import "fmt"

// ListNode ...
// type stack struct {
// 	vals []int
// }

func isValidSudoku(board [][]byte) bool {
	var pos = [4]uint64{0, 0, 0, 0}
	var blockNo, blockPos, hPos, vPos int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := int(board[i][j] - 48)
			if num+2 == 0 {
				continue
			}

			blockNo = (i/3)*3 + j/3
			blockPos = num + blockNo*9
			vPos = blockPos + 81
			hPos = vPos + 81

			if pos[blockPos/64]>>(blockPos%64)&1 == 1 ||
				pos[vPos/64]>>(vPos%64)&1 == 1 ||
				pos[hPos/64]>>(hPos%64)&1 == 1 {
				return false
			}
			pos[blockPos/64] += 0x01 << (blockPos % 64)
			pos[vPos/64] += 0x01 << (vPos % 64)
			pos[hPos/64] += 0x01 << (hPos % 64)

		}
	}
	return true
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
