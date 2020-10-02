Given a string s, find the length of the longest substring without repeating characters.

Example 1:
```
Input: s = "abcabcbb"
Output: 3
```
Explanation: The answer is "abc", with the length of 3.
Example 2:
```
Input: s = "bbbbb"
Output: 1
```
Explanation: The answer is "b", with the length of 1.
Example 3:
```
Input: s = "pwwkew"
Output: 3
```
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
Example 4:
```
Input: s = ""
Output: 0
``` 

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.

思路1
1. 需要3个变量，分别记录子串起始位置，最大子串长度和用切片保存子串
2. 每遍历到下一个字母，都需要遍历子串看是否有相同字母
3. 如果有相同字母，需要更新子串起始位置（原起始位置+冲突字母在现子串的位置顺序）
优点：切片保存子串信息，占用内存少
缺点：遍历子串查冲突的方法在遇到长子串时效率会低
   
```go
// runtime 4ms faster 91.75% memory 2.6MB less than 86.63%
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	st, maxLen := 0, 1
	curSub := s[0 : st+1]
	for i, n := range s[1:] {
		for j, m := range curSub {
			if n == m {
				st = st + j + 1
				curSub = s[st : i+2]
				continue
			}
		}
		curSub = s[st : i+2]
		if len(curSub) > maxLen {
			maxLen = len(curSub)
		}
	}
	return maxLen
}
```

思路2
与思路1的区别
1. 使用map保存子串位置，构成map[rune]int格式
2. 查冲突的方式变成直接ok:=map["word"]一次就能得到结果
优点：查冲突速度提升
缺点；需要额外的map内存记录

```go
//runtime 4ms memory 3.3MB
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	record := map[rune]int{rune(s[0]): 1}
	startPos, curLen, maxLen := 1, 1, 1
	for i, n := range s[1:] {
		if p, ok := record[n]; ok && p >= startPos {
			startPos = p + 1
			curLen = i - startPos + 3
			record[n] = i + 2
			continue
		}
		record[n] = i + 2
		curLen++
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
```