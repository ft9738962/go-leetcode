package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	var left, right, step, maxLen = 0, 0, 0, 1
	var maxSub = runes[0:1]
	for right < len(runes)-1 {
		for _, r := range runes[(left + 1):] {
			if runes[left] != r {
				break
			}
			right++
		}
		for i := 0; i < min(left, right, len(runes)); i++ {
			if runes[left-i-1] != runes[right+i+1] {
				break
			}
			step++
		}
		if right-left+2*step+1 > maxLen {
			maxLen = right - left + 2*step + 1
			maxSub = runes[(left - step):(right + step + 1)]
		}
		left = right + 1
		right = left
		step = 0
	}
	return string(maxSub)
}

func min(left, right, length int) int {
	if left > (length - right - 1) {
		return length - right - 1
	}
	return left
}

func main() {
	test1 := "a"
	test2 := "我家家我是非非二区路"
	test3 := "cbbd"

	fmt.Println(longestPalindrome(test1))
	fmt.Println(longestPalindrome(test2))
	fmt.Println(longestPalindrome(test3))
}
