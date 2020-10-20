Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

**Example 1:**
```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```
**Example 2:**
```
Input: n = 1
Output: ["()"]
```

**Constraints:**

1 <= n <= 8

### 思路:
采用递归方法，有以下递归调用规则：
1. 终止条件：左括号和右括号的数量等于n值
2. 发展条件：
   1. 当右括号数量小于左括号，且左括号数量小于n值，字符串可以加左括号或右括号并递归调用
   2. 当左括号等于n值，则只能递归调用添加右括号
   3. 当左括号数量等于右括号，则只能递归调用添加左括号

```go
// runtime 0ms 100% memory 2.8MB 6.22%
func generateParenthesis(n int) []string {

	result := []string{}
	recursiveGenerate(&result, "", n, 0, 0, 0)
	return result
}

func recursiveGenerate(result *[]string, current string,
	length, curLength, left, right int) {
	if length == left && length == right {
		*result = append(*result, current)
	} else if left < length && left > right {
		recursiveGenerate(result, current+"(",
			length, curLength+1, left+1, right)
		recursiveGenerate(result, current+")",
			length, curLength+1, left, right+1)
	} else if left == length {
		recursiveGenerate(result, current+")",
			length, curLength+1, left, right+1)
	} else {
		recursiveGenerate(result, current+"(",
			length, curLength+1, left+1, right)
	}
}
```