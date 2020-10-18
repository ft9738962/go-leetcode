package main

// ListNode ...
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	if string(s[len(s)-1]) == "(" ||
		string(s[len(s)-1]) == "[" ||
		string(s[len(s)-1]) == "{" {
		return false
	}

	var left string
	var loc int

	for len(s) > 0 {

		for i := len(s) - 2; i >= 0; i-- {
			if string(s[i]) == "(" ||
				string(s[i]) == "[" ||
				string(s[i]) == "{" {
				left = string(s[i])
				loc = i
				break
			} else if i == 0 {
				return false
			}
		}
		for j := loc + 1; j < len(s); j = j + 2 {
			pair := (left == "(" && string(s[j]) == ")") ||
				(left == "[" && string(s[j]) == "]") ||
				(left == "{" && string(s[j]) == "}")
			if pair {
				if j == loc+1 {
					s = s[:loc] + s[loc+2:]
				} else {
					s = s[:loc] + s[loc+1:j] + s[j+1:]
				}
				break
			} else if j+2 >= len(s) {
				return false
			}
		}
	}
	return true
}

// func main() {
// 	// test1 := []int{1, 0, -1, 0, -2, 2}

// 	// rand.Seed(time.Now().UnixNano())
// 	// for i := 0; i < 100; i++ {
// 	// 	test1 = append(test1, rand.Intn(101)-50)
// }
