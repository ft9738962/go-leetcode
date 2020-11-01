Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

**Follow up:**

- Could you solve the problem in O(1) extra memory space?
- You may not alter the values in the list's nodes, only nodes itself may be changed.

**Example 1:**

```
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
```
**Example 2:**
```
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
```
**Example 3:**
```
Input: head = [1,2,3,4,5], k = 1
Output: [1,2,3,4,5]
```
**Example 4:**
```
Input: head = [1], k = 1
Output: [1]
```

**Constraints:**

The number of nodes in the list is in the range sz.
- 1 <= sz <= 5000
- 0 <= Node.val <= 1000
- 1 <= k <= sz

思路：
- 使用双指针，指针a从第2位遍历到最末位，寻找与前一位不同的数
- 当指针a找到了不同数时，指针b将这个数搬运到从第2位起，确保指针b前的数字都不同
- 最后返回指针b的后一位数字

```go
// runtime 8ms 88.5% memory4.6MB 100%
func removeDuplicates(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	nums = nums[:j]
	return j + 1
}
```