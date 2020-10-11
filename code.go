package main

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

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func main() {
}
