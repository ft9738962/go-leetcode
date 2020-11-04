Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.

**Example 1:**
```
Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]
```

**Example 2:**
```
Input: l1 = [], l2 = []
Output: []
```
**Example 3:**
```
Input: l1 = [], l2 = [0]
Output: [0]
```

**Constraints:**

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both l1 and l2 are sorted in non-decreasing order.

思路：
双指针分别比较l1和l2的值，取小者连接到结果数链，当取到某个指针的Next为nil时，说明这个数链已经取完，将另一个数链的剩余部分绑定在结果数链的Next上返回

```go
// runtime 0ms 100% 2.5MB 100%
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, cur *ListNode
	if l1 == head {
		return l2
	}
	if l2 == head {
		return l1
	}
	if comp(l1.Val, l2.Val) {
		head = l2
		cur = head

		if l2.Next == nil {
			cur.Next = l1
			return head
		}
		l2 = l2.Next
	} else {
		head = l1
		cur = head

		if l1.Next == nil {
			cur.Next = l2
			return head
		}
		l1 = l1.Next
	}

	for {
		if comp(l1.Val, l2.Val) {
			cur.Next = l2
			cur = cur.Next

			if l2.Next == nil {
				cur.Next = l1
				break
			}
			l2 = l2.Next
		} else {
			cur.Next = l1
			cur = cur.Next
			if l1.Next == nil {
				cur.Next = l2
				break
			}
			l1 = l1.Next
		}
	}

	return head
}

func comp(a, b int) bool {
	if a > b {
		return true
	}
	return false
}
```