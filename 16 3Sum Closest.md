Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.


Example 1:
```
Input: nums = [-1,2,1,-4], target = 1
Output: 2
```
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 

Constraints:

- 3 <= nums.length <= 10^3
- -10^3 <= nums[i] <= 10^3
- -10^4 <= target <= 10^4

### 思路：

思路类似15题

先排序，后续使用三指针，外圈循环指针从小到大，循环前n-2个数（n为nums内数的个数），内圈双指针，通过计算与target的差，来调整双指针的逼近路线

```go
// runtime 4ms 97.95% memory 2.7MB 100%
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var gap = abs(target - nums[0] - nums[1] - nums[len(nums)-1])
	var output = nums[0] + nums[1] + nums[len(nums)-1]
	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			dist := target - nums[i] - nums[j] - nums[k]
			fmt.Println("g ", g)
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
```