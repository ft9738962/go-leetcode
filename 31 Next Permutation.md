Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).

The replacement must be in place and use only constant extra memory.

 

**Example 1:**
```
Input: nums = [1,2,3]
Output: [1,3,2]
```
**Example 2:**
```
Input: nums = [3,2,1]
Output: [1,2,3]
```
**Example 3:**
```
Input: nums = [1,1,5]
Output: [1,5,1]
```
**Example 4:**
```
Input: nums = [1]
Output: [1]
```

### Constraints:

- 1 <= nums.length <= 100
- 0 <= nums[i] <= 100

### 思路：
- 倒序寻找前后顺排的（因为全部倒排是字典最大，所有出现顺排说明还没有字典最大）
- 将找到的顺排前值，与后面所有的值比较，找到合理的插入位置，与其位置互换
- 将互换后的倒序排列变成顺排

```go
// runtime 0ms 100% 2.5MB 15.81%
func nextPermutation(nums []int) {
	var a, b int
	a, b = 0, len(nums)-1

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			if nums[i] < nums[len(nums)-1] {
				nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
				a = i + 1
				break
			}
			if i+2 <= len(nums)-1 {
				for j := i + 2; j <= len(nums)-1; j++ {
					if nums[i] >= nums[j] {
						nums[i], nums[j-1] = nums[j-1], nums[i]
						break
					}
				}
			} else {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
			a = i + 1
			break
		}
	}

	for a < b {
		nums[a], nums[b] = nums[b], nums[a]
		a++
		b--
	}
}
```