Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

Follow up: The overall run time complexity should be O(log (m+n)).

 

Example 1:
```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
```
Explanation: merged array = [1,2,3] and median is 2.

Example 2:
```
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
```
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

Example 3:
```
Input: nums1 = [0,0], nums2 = [0,0]
Output: 0.00000
```
Example 4:
```
Input: nums1 = [], nums2 = [1]
Output: 1.00000
```
Example 5:
```
Input: nums1 = [2], nums2 = []
Output: 2.00000
```

思路：
1. 因为题目要求复杂度为O(log(m+n))，因此不能使用遍历，考虑使用二分法。
2. 对短的切片进行两分，设这个分割点的左半边长度为x，同时因为两个切片长度已知，可以得到中位数的左右侧数组长度z（两数组长度和的一半），因此可以固定长数组对应分割点的左半边长度为y
3. 最终的目标就是在短切片中找到一个分割点，使得短切片左边x个数和长切片左边y个数都小于长短切片分割点的右边
4. 最后根据z是单数还是双数，决定怎么拿两个切片分割点左右的4个数求得中位数


```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, hi, midSize, mid1, mid2 := 0, m, (m+n+1)>>1, 0, 0
	for low <= hi {
		mid1 = low + (hi-low)>>1
		mid2 = midSize - mid1
		if mid1 != 0 && nums1[mid1-1] > nums2[mid2] {
			hi = mid1 - 1
		} else if mid1 != len(nums1) && nums1[mid1] < nums2[mid2-1] {
			low = mid1 + 1
		} else {
			break
		}
	}

	midLeft, midRight := 0, 0
	switch {
	case mid1 == 0:
		midLeft = nums2[mid2-1]
	case mid2 == 0:
		midLeft = nums1[mid1-1]
	default:
		_, midLeft = order(nums1[mid1-1], nums2[mid2-1])
	}
	if (m+n)&1 == 1 {
		return float64(midLeft)
	}

	switch {
	case mid1 == m:
		midRight = nums2[mid2]
	case mid2 == n:
		midRight = nums1[mid1]
	default:
		midRight, _ = order(nums1[mid1], nums2[mid2])
	}
	return float64(midLeft+midRight) / 2
}

func order(num1, num2 int) (int, int) {
	if num1 > num2 {
		return num2, num1
	}
	return num1, num2
}
```