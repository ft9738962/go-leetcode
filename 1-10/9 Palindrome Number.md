Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Follow up: Could you solve it without converting the integer to a string?

 

Example 1:
```
Input: x = 121
Output: true
```
Example 2:
```
Input: x = -121
Output: false
```
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:
```
Input: x = 10
Output: false
```
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
Example 4:
```
Input: x = -101
Output: false
```

Constraints:

-2^31 <= x <= 2^31 - 1

思路1：
把数字变成数组，数组对称比较，缺点构造新数组需要占用内存，比较数组的值会慢
```go
//runtime 24ms memory 6MB
func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}

	var nums []int
	for x != 0 {
		nums = append(nums, x%10)
		x = x / 10
	}

	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}
	return true
}
```

思路2：
构造一个倒转函数（把数字倒转过来），与原数字比较，这个方法速度快，更好
```go
 // runtime 8ms memory 5.2MB
 func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}

	return x == reverse(x)
}

// 倒转函数
func reverse(x int) int {
	var y int

	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}
	return y
}
```
