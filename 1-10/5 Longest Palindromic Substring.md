Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:
```
Input: "babad"
Output: "bab"
```
Note: "aba" is also a valid answer.
Example 2:
```
Input: "cbbd"
Output: "bb"
```

思路：



```go
// runtime 0ms 100% memory 2.3MB 50% 
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	bs := []byte(s)
	var left, right, step, maxLen = 0, 0, 0, 1
	var maxSub = bs[0:1]
	for right < len(s)-1 {
		for _, r := range bs[(left + 1):] {
			if s[left] != r {
				break
			}
			right++
		}
		for i := 0; i < min(left, right, len(bs)); i++ {
			if s[left-i-1] != s[right+i+1] {
				break
			}
			step++
		}
		if right-left+2*step+1 > maxLen {
			maxLen = right - left + 2*step + 1
			maxSub = bs[(left - step):(right + step + 1)]
		}
		left = right + 1
		right = left
		step = 0
	}
	return string(maxSub)
}

func min(left, right, length int) int {
	if left > (length - right -1) {
		return length - right - 1
	}
	return left
}
```