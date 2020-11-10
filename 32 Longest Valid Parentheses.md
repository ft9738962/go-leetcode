Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

 

**Example 1:**
```
Input: s = "(()"
Output: 2
```
Explanation: The longest valid parentheses substring is "()".

**Example 2:**
```
Input: s = ")()())"
Output: 4
```
Explanation: The longest valid parentheses substring is "()()".

**Example 3:**
```
Input: s = ""
Output: 0
```

### Constraints:

- 0 <= s.length <= 3 * 104
- s[i] is '(', or ')'.

### 思路1：
使用DP法，用数字切片记录每一位的当前的括号长度
有4种情况判断：
1. 如果当前是左括号，当前位长度为0
2. 如果当前是右括号，前一位是左括号，则长度是前2位长度+2
3. 如果当前是右括号，前一位是右括号，当前位减去前一位的长度再减1的位置是左括号（相当于"((()))"，第6位找第1位是不是左括号），则长度需要看当前位减去前一位的长度再减2的位置如果存在，长度等于这段的长度加上前一位长度再加2；如果不存在，就是前一位长度再加2
4. 其它情况都是长度为0
最后遍历切片寻找最大值

```go
// runtime 0ms 100% memory 4.9MB 7.26%
func longestValidParentheses(s string) int {
	var longest []int
	longest = []int{0}
	var tmp int
	for i := 1; i < len(s); i++ {
		if string(s[i]) == "(" {
			longest = append(longest, 0)
		} else {
			if i-1 >= 0 && string(s[i-1]) == "(" {
				if i == 1 {
					longest = append(longest, 2)
				} else {
					longest = append(longest, longest[i-2]+2)
				}
			} else if i-1 >= 0 && string(s[i-1]) == ")" && i-longest[i-1]-1 >= 0 && string(s[i-longest[i-1]-1]) == "(" {
				if i-longest[i-1]-2 >= 0 {
					tmp = longest[i-longest[i-1]-2]
				} else {
					tmp = 0
				}
				longest = append(longest, longest[i-1]+2+tmp)
			} else {
				longest = append(longest, 0)
			}
		}
	}
	return maxLongest(longest)
}

func maxLongest(longest []int) int {
	maxLen := 0
	for i := 0; i < len(longest); i++ {
		if longest[i] > maxLen {
			maxLen = longest[i]
		}
	}
	return maxLen
}
```

### 思路2：
使用栈来储存妨碍括号连续的括号位置，这样只要遍历一遍后，就可以获得所有中断点的位置，每个中断点之间的长度就是连续括号的长度，再遍历一遍后就能得到最大值，栈的规则如下：
- 如果是左括号，编号入栈
- 如果是右括号，且栈内最上是左括号，则弹出这个左括号的编号
- 如果是右括号，但不符合上一条，编号入栈
- 遍历后在栈最下插入-1，最上插入字符串长度
- 再次遍历栈内所有相邻编号，得到最大长度

```go
// runtime 0ms 100% memory 3.8MB 7.25%
type stack struct {
	vals []int
}

func (s *stack) push(n int) {
	s.vals = append(s.vals, n)
}

func (s *stack) pop() {
	s.vals = s.vals[:len(s.vals)-1]
}

func (s *stack) insert(pos int, n int) {
	if pos < 0 || pos > len(s.vals) {
		fmt.Println("position is not valid")
	}
	if pos == 0 {
		s.vals = append([]int{n}, s.vals...)
	} else if pos == len(s.vals) {
		s.push(n)
	} else {
		s.vals = append(append(s.vals[:pos], n), s.vals[pos+1:]...)
	}
}

func (s *stack) peak() int {
	return s.vals[len(s.vals)-1]
}

func longestValidParentheses(s string) int {
	st := stack{[]int{}}
	var maxLen int
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			st.push(i)
		} else {
			if len(st.vals) > 0 {
				if string(s[st.peak()]) == "(" {
					st.pop()
				} else {
					st.push(i)
				}
			} else {
				st.push(i)
			}
		}
	}
	st.insert(0, -1)
	st.push(len(s))
	for i := 0; i < len(st.vals)-1; i++ {
		if st.vals[i+1]-st.vals[i]-1 > maxLen {
			maxLen = st.vals[i+1] - st.vals[i] - 1
		}
	}
	return maxLen
}
```