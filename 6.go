package main

import (
	"fmt"
	"strings"
)

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

func main() {
	s1 := "古埃及打仗用的投掷武器和箭，打完了还回收吗？这些武器涂毒吗？回收的时候怎么区分哪些有毒哪些没毒呢？"
	s2 := "ABC"
	s3 := "ABCDEFGHIJKLMNOPQ"
	n1 := 3
	n2 := 4
	fmt.Print(convert(s2, n1))
	fmt.Println()
	fmt.Print(convert(s1, n2))
	fmt.Println()
	fmt.Println(convert(s3, n2))
}
