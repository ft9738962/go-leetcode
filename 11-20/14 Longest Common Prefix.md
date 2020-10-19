Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

 

Example 1:
```
Input: strs = ["flower","flow","flight"]
Output: "fl"
```
Example 2:
```
Input: strs = ["dog","racecar","car"]
Output: ""
```
Explanation: There is no common prefix among the input strings.
 

Constraints:

0 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lower-case English letters.

思路：
使用一个变量保存公共前置，然后与strs内的单词逐个比对，如果新的公共前置词比变量短，就更新变量

```go
// runtime 0ms 100% memory 2.4MB 7.72% 
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	cur := strs[0]

	for _, word := range strs[1:] {
        if len(cur) == 0 {
            return ""
        } else if len(cur) > len(word) {
			cur, word = word, cur
		}
		for i := 0; i < len(cur); i++ {
			if cur[i] == word[i] {
				continue
			}
			cur = cur[:i]
			break
		}
	}
	return cur
}
```