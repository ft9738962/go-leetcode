Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
 

**Example 1:**
```
Input: s = "()"
Output: true
```
**Example 2:**
```
Input: s = "()[]{}"
Output: true
```
**Example 3:**
```
Input: s = "(]"
Output: false
```
**Example 4:**
```
Input: s = "([)]"
Output: false
```
**Example 5:**
```
Input: s = "{[]}"
Output: true
```

思路1：
使用字典区分左边符号和右边符号，左边符号放入字符串切片，出现右边符号就与切片的末尾的左边符号比较，如果对应，就把切片最后的左边符号抛弃。这样遍历结束后，查看字符串是否清空，如果没有，就说明有左边符号没有对上，返回错误
```go
// runtime 0ms 100% memory 2.2MB 6.89%
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	validMap := map[string]string{
		"(": "",
		"[": "",
		"{": "",
		")": "(",
		"]": "[",
		"}": "{",
	}

	left := []string{}
	for _, r := range s {

		if validMap[string(r)] == "" {
			left = append(left, string(r))
		} else if len(left) >= 1 {
			if validMap[string(r)] ==
				left[len(left)-1] {
				left = left[:len(left)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if len(left) > 0 {
		return false
	}
	return true
}
```

思路2：
将字典变成ascii码硬编码，可以减少内存使用
```go
// runtime 0ms 100% 2MB 6.89%
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	left := []int32{}
	for _, r := range s {

		if r == 40 || r == 91 || r == 123 {
			left = append(left, r)
		} else if len(left) >= 1 {
			if r == 41 && left[len(left)-1] == 40 {
				left = left[:len(left)-1]
			} else if r == 93 && left[len(left)-1] == 91 {
				left = left[:len(left)-1]
			} else if r == 125 && left[len(left)-1] == 123 {
				left = left[:len(left)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if len(left) > 0 {
		return false
	}
	return true
}
```