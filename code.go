package main

import (
	"fmt"
	"sort"
)

type hashSet struct {
	hash   map[string]bool
	result [][]int
}

func newHashSet() *hashSet {
	return &hashSet{
		make(map[string]bool),
		make([][]int, 0),
	}
}

func (hs *hashSet) add(arr []int) {
	if hs.hash[fmt.Sprint(arr)] == false {
		hs.hash[fmt.Sprint(arr)] = true
		hs.result = append(hs.result,
			arr)
	}
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums) // O(nlog(n))
	var hs = newHashSet()

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 {
			if nums[i-1]+nums[i] > 0 {
				break
			}
		}

		for j, k := i+1, len(nums)-1; j < k; {
			if nums[i]+nums[j]+nums[k] == 0 {
				hs.add([]int{nums[i], nums[j], nums[k]})
				j++
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}

		}
	}
	return hs.result
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func main() {
	// test1 := make([]int, 0)
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 100; i++ {
	// 	test1 = append(test1, rand.Intn(101)-50)
	// }
	test1 := []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum(test1))
}
