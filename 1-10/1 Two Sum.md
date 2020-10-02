## 1. Two Sum ##

Given an array of integers nums and an integer target, 
return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, 
and you may not use the same element twice.
You can return the answer in any order.

Example 1:
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

Example 2:
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```
Example 3:
```
Input: nums = [3,3], target = 6
Output: [0,1]
```

the input may include negative number
since there is only one solution, if one pair of number is found, 
the result can be returned 
method1: save memory complexity is O(n^2)
```go
func twoSum(nums []int, target int) []int {
    for i, n:= range nums {
		for j, m := range nums[i+1:] {
			if n + m == target {
				return []int{i, i+j+1}
			}
		}
	}
	return nil
}
```

use a map recording left number to save time
```go
func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
    for i, n:= range nums {
		if m, ok := temp[n]; ok {
			return []int{m,i}
		}
        temp[target-n] = i
	}
	return nil
}
```