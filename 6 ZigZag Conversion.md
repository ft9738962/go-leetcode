The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
```
P   A   H   N
A P L S I I G
Y   I   R
```
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:
```
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
```
Example 2:
```
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
```
Explanation:
```
P     I    N
A   L S  I G
Y A   H R
P     I
```

思路：
1. 直接按照题目输出要求接，根据numRows的规则计算每一行需要输出的string
```go
```
2. 先构筑ZigZag数组，可以直接输出ZigZag字符，然后再把空格去掉拼接成输出
```go
// convert主体函数，控制整个流程，将string先转换成[]rune（这样可以对中文使用），然后按照行数要求，分成若干段子rune切片
// 把每个rune切片交付groupConvert转化成每行按照rune和空格排列的rune切片的切片[][]rune
// 由groupCombine把每个[][]rune合并
// 最后由groupFormatter输出成要求的格式
func convert(s string, numRows int) string {
	var (
		st, ed            int
		newGroups, groups [][]rune
	)

	if len(s) == 0 {
		return ""
	} else if numRows == 1 {
		return s
	}
	r := []rune(s)
	st, ed = 0, 2*numRows-2
	for st < len(r) {
		if ed > len(r) {
			ed = len(r)
		}
		newGroups = groupConvert(r[st:ed], numRows)
		groups = groupCombine(groups, newGroups)
		st = ed
		ed += 2*numRows - 2
	}
	return groupFormater(groups)
}

func groupConvert(r []rune, numRows int) [][]rune {
	var (
		rowLength, step int
		initLine        []rune
		group           [][]rune
	)

	if len(r) <= numRows {
		rowLength = 1
	} else {
		rowLength = len(r) - numRows + 1
		step = numRows*2 - len(r) - 2
	}
	for i := 0; i < min(numRows, len(r)); i++ {
		initLine = []rune(strings.Repeat(" ", rowLength))
		initLine[0] = r[i]
		if i > step && rowLength > 1 && i != numRows-1 {
			initLine[rowLength-i+step] = r[len(r)-i+step]
		}
		group = append(group, initLine)
	}
	return group
}

func groupCombine(formerGroups, newGroups [][]rune) [][]rune {
	if len(formerGroups) == 0 {
		return newGroups
	}
	for i, line := range newGroups {
		formerGroups[i] = append(formerGroups[i], line...)
	}
	return formerGroups
}

func groupFormater(r [][]rune) (s string) {
	for _, line := range r {
		s = s + strings.ReplaceAll(string(line), " ", "")
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```