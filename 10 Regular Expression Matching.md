Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*' where: 

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

 

Example 1:
```
Input: s = "aa", p = "a"
Output: false
```
Explanation: "a" does not match the entire string "aa".
Example 2:
```
Input: s = "aa", p = "a*"
Output: true
```
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:
```
Input: s = "ab", p = ".*"
Output: true
```
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:
```
Input: s = "aab", p = "c*a*b"
Output: true
```
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
Example 5:
```
Input: s = "mississippi", p = "mis*is*p*."
Output: false
```

Constraints:
- 0 <= s.length <= 20
- 0 <= p.length <= 30
- s contains only lowercase English letters.
- p contains only lowercase English letters, '.', and '*'.

思路：
动态规划法：
假设s和p字符串的长度为m和n，构造一个(m+1)*(n+1)的[][]bool数组用来储存dp的结果

分别用i和j对s和p从位置0和位置1进行遍历，有以下规则：
- 位置0表示空位，位置i表示s\[i-1]，位置j表示p\[j-1]
- dp\[0]\[0]为true，表示两个字符串都是空时，结果为true
- 当位置j对应的p\[j-1]不是"\*"时，判断以下三者是否都为真：
  - i是否在空位(0)
  - s\[i-1]和p\[j-1]是否相同，或者p\[j-1]是否为"."
  - dp\[i-1]\[j-1]是否为true
- 当位置j对应的p\[j-1]是"\*"时，对以下这个条件进行“或”判断：
  - dp\[i]\[j-2]是否为真（对应"aa"和"ab*a"情况）
  - 当i大于1时，s\[i-1]和p\[j-2]是否相同，或者p\[j-2]是否为"."，且dp\[i-1]\[j]是否为真（对应"aaa"和"a*"情况）

```go
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if j > 1 && p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j]
			} else {
				dp[i][j] = i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') && dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
```