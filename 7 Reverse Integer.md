Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
```
Input: 123
Output: 321
```
Example 2:
```
Input: -123
Output: -321
```
Example 3:
```
Input: 120
Output: 21
```
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

思路：
翻转数字就是不停的求10余，将余数倒转顺序乘以10

```go
func reverse(x int) int {
	if x == 0 {
		return 0
	}
	var r int

	for x != 0 {
		r = x%10 + r*10
		x = x / 10
	}
	if r > 1<<31-1 || r < -(1<<31) {
		return 0
	}
	return r

}
```