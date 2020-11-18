Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

Follow up: Could you write an algorithm with O(log n) runtime complexity?

 

**Example 1:**
```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```
**Example 2:**
```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```
**Example 3:**
```
Input: nums = [], target = 0
Output: [-1,-1]
```

**Constraints:**

- 0 <= nums.length <= $10^5$
- $-10^9$ <= nums[i] <= $10^9$
- nums is a non-decreasing array.
- $-10^9$ <= target <= $10^9$

### 思路1：
1. 先找到1个目标值，然后往左右两边延伸，寻找左右边界
```go
// runtime 8ms 79.87% memory 4.1MB 100%
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l, mid, r := 0, (len(nums)-1)/2, len(nums)-1
	found := 0
	if nums[len(nums)-1] == target {
		found = 1
		mid = r
	} else if nums[0] == target {
		found = 1
		mid = l
    }
    // 找一个符合的目标值位置
	for l != mid && found == 0 {
		if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		} else {
			found = 1
			break
		}
		mid = (l + r) / 2
    }
    
    // 找不到直接给返回
	if found == 0 {
		return []int{-1, -1}
	}

    // 找到后探寻边界
	lBound, rBound := mid, mid
	for lBound > 0 {
		lBound--
		if nums[lBound] != nums[mid] {
			lBound++
			break
		}
	}

	for rBound < len(nums)-1 {
		rBound++
		if nums[rBound] != nums[mid] {
			rBound--
			break
		}
	}

    // 返回边界
	return []int{lBound, rBound}
}
```

### 思路2：
1. 先找左边界位置（两分法）
2. 在左边界和末尾之间找右边界位置（两分法）
