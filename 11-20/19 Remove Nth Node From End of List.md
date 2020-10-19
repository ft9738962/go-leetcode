Given the head of a linked list, remove the $n^{th}$ node from the end of the list and return its head.

Follow up: Could you do this in one pass?

**Example 1:**
```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```
**Example 2:**
```
Input: head = [1], n = 1
Output: []
```
**Example 3:**
```
Input: head = [1,2], n = 1
Output: [1]
```

**Constraints:**

The number of nodes in the list is sz.
- 1 <= sz <= 30
- 0 <= Node.val <= 100
- 1 <= n <= sz

```go
// runtime 0ms 100% memory 2.2MB 100%
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	dist := 1
	var before = head
	for cur.Next != nil {
		if dist >= n+1 {
			before = before.Next
		}
		cur = cur.Next
		dist++
	}
	if dist == n {
		head = before.Next
	} else {
		before.Next = before.Next.Next
	}
	return head
}
```