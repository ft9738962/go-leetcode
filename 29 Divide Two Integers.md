Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.

**Note:**

Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For this problem, assume that your function returns 231 − 1 when the division result overflows.
 
**Example 1:**
```
Input: dividend = 10, divisor = 3
Output: 3
```
Explanation: 10/3 = truncate(3.33333..) = 3.
**Example 2:**
```
Input: dividend = 7, divisor = -3
Output: -2
```
Explanation: 7/-3 = truncate(-2.33333..) = -2.
**Example 3:**
```
Input: dividend = 0, divisor = 1
Output: 0
```
**Example 4:**
```
Input: dividend = 1, divisor = 1
Output: 1
```

**Constraints:**

- $-2^{31}$ <= dividend, divisor <= $2^{31}$ - 1
- divisor != 0

思路：
因为使用固定的divisor削减dividend太慢，使用指数方法快速增加削减速度，每次当削减速度超过dividend时，重置削减值
```go
// runtime 4ms 66.94% memory 2.5MB 17.14%
func divide(dividend int, divisor int) int {
	var sign, reducer, ct, result int
	sign = 1

	if dividend == 0 {
		return 0
	} else if divisor == 1 || divisor == -1 {
		result = dividend * divisor
	} else if dividend > 0 && divisor < 0 {
		sign = -1
		divisor = -1 * divisor
	} else if dividend < 0 && divisor > 0 {
		sign = -1
		dividend = -1 * dividend
	} else if dividend < 0 && divisor < 0 {
		dividend = -1 * dividend
		divisor = -1 * divisor
	}

	if result == 0 {
		ct = 1

		reducer = divisor
		for dividend >= divisor {
			if dividend < reducer {
				reducer = divisor
				ct = 1
			}
			dividend -= reducer
			result += ct
			reducer = reducer << 1
			ct = ct << 1
		}
	}

	if result > 1<<31-1 || result < -(1<<31) {
		return (1<<31 - 1) * sign
	}
	return result * sign
}
```