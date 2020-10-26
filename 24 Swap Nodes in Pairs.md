Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes. Only nodes itself may be changed.

**Example 1:**

```
Input: head = [1,2,3,4]
Output: [2,1,4,3]
```
**Example 2:**
```
Input: head = []
Output: []
```
**Example 3:**
```
Input: head = [1]
Output: [1]
```

**Constraints:**

- The number of nodes in the list is in the range [0, 100].
- 0 <= Node.val <= 100

### 思路：
三指针遍历，一个指针a指向前一组交换后的后者，另外两个指针b和c指向需要交换的两个Listnode：
- a指向c
- c指向b
- b指向原来c的Next

```go
// runtime 0ms 100% memory 2.1MB 100%
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	a := head
	b := head.Next
	head = b
	var last *ListNode
	last = nil

	for {
		if last == nil {
			last = a
		} else {
			last.Next = b
		}
		a.Next, b.Next = b.Next, a
		if a.Next == nil {
			break
		}
		if a.Next.Next == nil {
			break
		}
		last, a, b = a, a.Next, a.Next.Next
	}

	return head
}
```