Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.

**Example 1:**
```
Input: board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
```
```
Output: true
```
**Example 2:**
```
Input: board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
```
```
Output: false
```
**Explanation:** 

Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
**Constraints:**

- board.length == 9
- board[i].length == 9
- board[i][j] is a digit or '.'.

### 思路1：
这题的思路主要是用怎么样最小的数据格式储存已有的数独信息，第一种方法是效率较低的切片+字典形式储存，需要使用根据循环的进度，每9个数字各使用一个字典确认数字是否重复

```go
// runtime 4ms 70.33% memory 3.6MB 44.72%
func isValidSudoku(board [][]byte) bool {
	var hlineCheck map[byte]bool
	var vlineCheck map[byte]bool
	blockCheck := [9]map[byte]bool{}
	for a := 0; a < 9; a++ {
		blockCheck[a] = map[byte]bool{}
	}

	for i := 0; i < 9; i++ {
		hlineCheck = map[byte]bool{}
		vlineCheck = map[byte]bool{}
		for j := 0; j < 9; j++ {
			if board[i][j] == 46 {
			} else if hlineCheck[board[i][j]] == false {
				hlineCheck[board[i][j]] = true
			} else {
				return false
			}

			if board[j][i] == 46 {
			} else if vlineCheck[board[j][i]] == false {
				vlineCheck[board[j][i]] = true
			} else {
				return false
			}

			if board[i][j] == 46 {
			} else if blockCheck[(i/3)*3+j/3][board[i][j]] == false {
				blockCheck[(i/3)*3+j/3][board[i][j]] = true
			} else {
				return false
			}
		}
	}
	return true
}
```

### 思路2
使用bit来储存，用9个9bit来表示横排，9个9bit表示竖排，9个9bit表示区块，这样只需要243个bit就能够包含数独的所有信息，对应的就是4个uint64数据

```go
// runtime 0ms 100% memory 2.8MB 100%
func isValidSudoku(board [][]byte) bool {
	var pos = [4]uint64{0, 0, 0, 0} // 初始4个uint64数字储存
	var blockNo, blockPos, hPos, vPos int // 记录数字在行、列、块的位置
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := int(board[i][j]) - 49
			if num+3 == 0 {
				continue
			} // 跳过"."

			blockNo = (i/3)*3 + j/3
			blockPos = num + blockNo*9
			vPos = 81 + j*9 + num
			hPos = 162 + i*9 + num

			if pos[blockPos/64]>>(blockPos%64)&1 == 1 ||
				pos[vPos/64]>>(vPos%64)&1 == 1 ||
				pos[hPos/64]>>(hPos%64)&1 == 1 {
				return false
			} // 如果该数字在同行、同列、同块有重复，返回false
			pos[blockPos/64] += 0x1 << (blockPos % 64)
			pos[vPos/64] += 0x1 << (vPos % 64)
            pos[hPos/64] += 0x1 << (hPos % 64)
            // 没有的话，就在对应位置记录

		}
	}
	return true
}
```