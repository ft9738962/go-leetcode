You are given a string s and an array of strings words of the same length. Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.

You can return the answer in any order.

 

**Example 1:**
```
Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0,9]
```
Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.

**Example 2:**
```
Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
Output: []
```
**Example 3:**
```
Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
Output: [6,9,12]
```

**Constraints:**

- 1 <= s.length <= 104
- s consists of lower-case English letters.
- 1 <= words.length <= 5000
- 1 <= words[i].length <= 30
- words[i] consists of lower-case English letters.

### 思路：
- 建立一个需要单词哈希表，遍历单词数组，将同一个词出现的次数存放在哈希表中
- 按照子词组的长度逐一遍历字符串，将可能的子字符串按照单词的长度，分别与哈希表中对照，如果没有这个词就跳至下一个子字符串
- 建立一个临时的子字符串哈希表，保存每个单词的出现次数，将出现次数和原哈希表的出现次数比对，出现次数超过就跳至下一个子字符串
- 如果全部符合，就把这个子字符串头部的索引顺序值存入结果数组中

```go
// runtime 36ms 86.36% memory 5.7MB 12.5%
func findSubstring(s string, words []string) []int {
	wordNum := len(words)
	wordLen := len(words[0])
	subStringLen := wordNum * wordLen
	result := make([]int, 0)
	var pSubstring string

	wordsMap := map[string]int{}
	for _, word := range words {
		_, ok := wordsMap[word]
		if ok == false {
			wordsMap[word] = 1
		} else {
			wordsMap[word]++
		}
	}

	for i := 0; i < len(s)+1-subStringLen; i++ {
		pSubstring = s[i : i+subStringLen]
		wordsCount := map[string]int{}
		valid := 1
		for j := 0; j < subStringLen; j = j + wordLen {
			w := pSubstring[j : j+wordLen]
			num, ok := wordsMap[w]
			if ok == false {
				valid = 0
				break
			} else {
				ct, ok := wordsCount[w]
				if ok == false {
					wordsCount[w] = 1
				} else if ct == num {
					valid = 0
					break
				} else {
					wordsCount[w]++
				}
			}
		}
		if valid == 1 {
			result = append(result, i)
		}

	}
	return result
}
```