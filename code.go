package main

func romanToInt(s string) int {
	var result int

	for _, b := range s {
		switch b {
		case 'M':
			if result%1000 > 0 {
				result += 800
			} else {
				result += 1000
			}
		case 'D':
			if result%1000 > 0 {
				result += 300
			} else {
				result += 500
			}
		case 'C':
			if result%100 > 0 {
				result += 80
			} else {
				result += 100
			}
		case 'L':
			if result%100 > 0 {
				result += 30
			} else {
				result += 50
			}
		case 'X':
			if result%10 > 0 {
				result += 8
			} else {
				result += 10
			}
		case 'V':
			if result%10 > 0 {
				result += 3
			} else {
				result += 5
			}
		case 'I':
			result++
		}
	}
	return result
}

// func main() {
// 	fmt.Print(myAtoi("9223372036854775808"))
// }
