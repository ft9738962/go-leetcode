You are given an integer array nums sorted in ascending order, and an integer target.

Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

If target is found in the array return its index, otherwise, return -1.

 

**Example 1:**
```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```
**Example 2:**
```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```
**Example 3:**
```
Input: nums = [1], target = 0
Output: -1
```

**Constraints:**

- 1 <= nums.length <= 5000
- -10^4 <= nums[i] <= 10^4
- All values of nums are unique.
- nums is guranteed to be rotated at some pivot.
- -10^4 <= target <= 10^4

### 思路：
使用二分法锁定目标值的可能范围进行迭代，因为数组不是单向递增递减，且最左侧值，必定大于最右侧值，其中值有两种情况：
- 中值大于最左侧，从最左侧至中值为单向递增，中值至最右侧仍类似原来的情形
- 中值小于最右侧，从中值至最右侧为单向递增，最左侧至中值仍类似原来的情形
通过对目标值的迭代，找到目标值所属的单向递增区，然后使用普通的二分查找

```go
// runtime 0ms 100% memory 2.6MB 100%
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
```