package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result = make([][]int, 0)

	for i := range nums[:len(nums)-3] {
		if i > 1 {
			if nums[i] == nums[i-1] {
				continue
			}
		}
		rs := threeSum(nums[i+1:], target-nums[i])
		for _, r := range rs {
			var temp []int
			result = append(result, append(append(temp, nums[i]), r...))
		}
	}
	return result
}

func threeSum(nums []int, target int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums) // O(nlog(n))
	var numsDict = make(map[int][]int)
	var result [][]int

	for i, n := range nums {
		if _, ok := numsDict[n]; ok != true {
			numsDict[n] = []int{i}
		} else {
			numsDict[n] = append(numsDict[n], i)
		} // O(n)
	}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 {
			if nums[i-1] == nums[i] {
				continue
			}
			if nums[i-1]+nums[i] > 0 {
				break
			}
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[j-1] == nums[j] && j > i+1 {
				continue
			}
			if locs, ok := numsDict[target-nums[i]-nums[j]]; ok == true {
				if locs[len(locs)-1] <= j {
					continue
				}
				result = append(result,
					[]int{nums[i], nums[j], target - nums[i] - nums[j]})
			}
			if nums[i]+nums[j]-target > 0 {
				break
			}
		}
	}
	return result
}

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }

func main() {
	test1 := make([]int, 0)
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 100; i++ {
	// 	test1 = append(test1, rand.Intn(101)-50)
	// }
	fmt.Println(fourSum(test1, 1))
}
