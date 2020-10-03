package main

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

// func main() {
// 	fmt.Print(myAtoi("9223372036854775808"))
// }
