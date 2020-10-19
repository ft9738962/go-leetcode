Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example 1:
```
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```
Example 2:
```
Input: digits = ""
Output: []
```
Example 3:
```
Input: digits = "2"
Output: ["a","b","c"]
```

Constraints:

- 0 <= digits.length <= 4
- digits[i] is a digit in the range ['2', '9'].

## 思路：
建立电话号码的间隔表和起点偏移表：
- 间隔表用来记录每个电话按键包含了几个字母
- 起点偏移表用来记录每个电话按键的首字母相对于字母a的间隔（比如d相对于a偏移为3，q相对于a偏移为15）

因为字符串2~9对应ascii码50~58，字母a~z对应ascii码97~112

所以字符串到字母的对应关系：
- 字符串数字-48=按键数字
- 按键数字的起点偏移+97+（字母在这个按键中的顺序-1）=字符串字母

然后使用一个字符串切片保存过程中的拼接结果
```go
// runtime 0ms 100% memory 2.1MB 100%
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var digInterval = []int{3, 3, 3, 3, 3, 4, 3, 4}
	var digStart = []int{0, 3, 6, 9, 12, 15, 19, 22}

	var result = make([]string, 0)
	for _, d := range digits {
		letters := make([]string, 0)
		dig := int(d) - 48
		for i := 0; i < digInterval[dig-50]; i++ {
			letters = append(letters, string(97+digStart[dig-2]+i))
		}
		if len(result) == 0 {
			result = letters
			continue
		}
		var temp = make([]string, 0)
		for _, char1 := range result {
			for _, char2 := range letters {
				temp = append(temp, char1+char2)
			}
		}
		result = temp
	}
	return result
}
```
