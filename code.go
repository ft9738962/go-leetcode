package main

import "fmt"

// ListNode ...
type stack struct {
	vals []int
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	result := -1
	recurSearch(nums, target, l, r, &result)
	return result
}

func recurSearch(nums []int, target int, l, r int, result *int) {
	mid := (l + r) / 2
	if mid == l {
		if target == nums[l] {
			*result = l
		}
		if target == nums[r] {
			*result = r
		}
	} else {
		if nums[mid] >= nums[l] {
			if target >= nums[l] {
				if nums[mid] >= target {
					oneDirectionSearch(nums, target, l, mid, result)
				} else {
					recurSearch(nums, target, mid+1, r, result)
				}
			} else {
				recurSearch(nums, target, mid+1, r, result)
			}
		} else {
			if target >= nums[l] {
				recurSearch(nums, target, l, mid, result)
			} else {
				if target <= nums[r] {
					if target >= nums[mid] {
						oneDirectionSearch(nums, target, mid, r, result)
					} else {
						recurSearch(nums, target, l, mid-1, result)
					}
				}
			}
		}
	}
}

func oneDirectionSearch(nums []int, target int, l, r int, result *int) {
	mid := (l + r) / 2
	for mid != l {
		if nums[r] == target {
			*result = r
			break
		} else if nums[mid] > target {
			r = mid
		} else if nums[mid] < target {
			l = mid
		} else {
			*result = mid
			break
		}
		mid = (l + r) / 2
	}
	if nums[l] == target {
		*result = l
	} else if nums[r] == target {
		*result = r
	}
}

func main() {

	// in := []int{4, 5, 6, 7, 0, 1, 2, 3}
	in2 := []int{1, 3, 5}
	// in3 := ")()())"
	// in4 := ")(((((()())()()))()(()))("
	// in5 := ")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"
	// fmt.Println(longestValidParentheses(in))

	// fmt.Println(search(in, 0))
	fmt.Println(search(in2, 3))
	// fmt.Println(longestValidParentheses(in4))
	// fmt.Println(longestValidParentheses(in5))

	// fmt.Println(r2)

}
