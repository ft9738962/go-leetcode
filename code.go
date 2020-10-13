package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var gap = abs(target - nums[0] - nums[1] - nums[len(nums)-1])
	var output = nums[0] + nums[1] + nums[len(nums)-1]
	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			dist := target - nums[i] - nums[j] - nums[k]
			if dist == 0 {
				return nums[i] + nums[j] + nums[k]
			}
			if gap > abs(dist) {
				gap = abs(dist)
				output = nums[i] + nums[j] + nums[k]
			}
			if dist < 0 {
				k--
			} else {
				j++
			}
		}
	}
	fmt.Println(output)
	return output
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// func main() {
// 	// test1 := make([]int, 0)
// 	// rand.Seed(time.Now().UnixNano())
// 	// for i := 0; i < 100; i++ {
// 	// 	test1 = append(test1, rand.Intn(101)-50)
// 	// }
// 	test1 := []int{3, 0, -2, -1, 1, 2}
// 	fmt.Println(threeSum(test1))
// }
