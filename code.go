package main

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}

	return x == reverse(x)
}

func reverse(x int) int {
	var y int

	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}
	return y
}

// func main() {
// 	fmt.Print(myAtoi("9223372036854775808"))
// }
