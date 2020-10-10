Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
```
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```
For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

- I can be placed before V (5) and X (10) to make 4 and 9. 
- X can be placed before L (50) and C (100) to make 40 and 90. 
- C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral.

 

Example 1:
```
Input: num = 3
Output: "III"
```
Example 2:
```
Input: num = 4
Output: "IV"
```
Example 3:
```
Input: num = 9
Output: "IX"
```
Example 4:
```
Input: num = 58
Output: "LVIII"
```
Explanation: L = 50, V = 5, III = 3.

Example 5:
```
Input: num = 1994
Output: "MCMXCIV"
```
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
 

Constraints:

1 <= num <= 3999

思路1：
构造一个字典，将数字从大至小解析，寻找对应的数字规则输出罗马字符

这里使用了golang的string.Builder构造罗马数字，其对于内存的占用很小，比普通的string可以提高8ms
```go
// runtime 0ms 100% memory 5.4MB 3.36%
import "strings"

func intToRoman(num int) string {
	var romanMap = map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	var roman strings.Builder
	numlist := []int{1000, 900, 500, 400, 100,
		90, 50, 40, 10, 9, 5, 4, 1}

	for _, n := range numlist {
		for num-n >= 0 {
			roman.WriteString(romanMap[n])
			num -= n
		}
	}

	return roman.String()
}
```

思路2：
用数组保存不同的罗马数字，使用10进制的关系，每次取出不同位上的5和10的罗马字符，使用switch case来分析
```go
// runtime 12ms 53.85% memory 3.4MB 9.36%
func intToRoman(num int) string {
	n := 1000
	m := []string{"I", "V", "X", "L", "C", "D", "M", "", ""}
	result := ""
	s := 3
	for n > 0 {
		result += numToRoman(num/n, m[s*2], m[s*2+1], m[s*2+2])
		num -= num / n * n
		n = n / 10
		s--
	}
	return result
}

func numToRoman(num int, one string, five string, ten string) string {
	switch {
	case num < 4:
		return strings.Repeat(one, num)
	case num == 4:
		return one + five
	case num >= 5 && num < 9:
		return five + strings.Repeat(one, num-5)
	case num == 9:
		return one + ten
	default:
		return ""
	}
}
```