package main

func maxArea(height []int) int {
	var max, curArea int
	l, r := 0, len(height)-1
	for l < r {
		curArea = min(height[l], height[r]) * (r - l)
		if curArea > max {
			max = curArea
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Print(myAtoi("9223372036854775808"))
// }
