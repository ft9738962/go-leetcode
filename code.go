package main

import "fmt"

// ListNode ...
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

func main() {

	// in := ")()())()()()(())))()()()()"
	// in2 := "()(())"
	in3 := ")()())"
	in4 := ")(((((()())()()))()(()))("
	in5 := ")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"
	// fmt.Println(longestValidParentheses(in))

	// fmt.Println(longestValidParentheses(in2))
	fmt.Println(longestValidParentheses(in3))
	fmt.Println(longestValidParentheses(in4))
	fmt.Println(longestValidParentheses(in5))

	// fmt.Println(r2)

}
