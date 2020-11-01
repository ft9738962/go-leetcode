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

### 思路
题意要求空间复杂度O(1)，必须使用固定数量的变量来进行反转，将这个过程进行拆分：
- 将左->右反转成左<-右需要三个变量记录：左，右，右的右，操作顺序为：
  1. 右->左
  2. 右的右赋值给右；右赋值给左
  3. 生成新的右的右
- 对于反转一组k个值需要两个变量：k个节点的首个节点first，first的上一个节点prev，在反转完所有值后，其操作为：
    1. first->right（right此时已经变成了下一组k个节点的首个节点，first是k节点的首个，反转后变成最后一个，所以要指向下一组节点的首节点）
    2. prev->left(left此时是k节点的最后一个，反转后变成首个节点，所以要将prev指向它)
    3. 更新prev和first的值（将原first赋值给prev，然后将prev的下个节点赋值给first）

```go
// runtime 4ms 98.77%  memory 3.6MB 100%
func reverseKGroup(head *ListNode, k int) *ListNode {

	if head.Next == nil {
		return head
	}
	ct := countListNodes(head)
	temp := &ListNode{Val: -1, Next: head}
	head = temp
	prev := head
	var left, right, next, first *ListNode
	for i := 0; i < ct/k; i++ {
		left = prev.Next
		right = left.Next
		for j := 0; j < k-1; j++ {
			next = right.Next
			right.Next = left
			left = right
			right = next
		}
		prev.Next.Next = right
		first = prev.Next
		prev.Next = left
		prev = first
	}

	return head.Next
}

func countListNodes(head *ListNode) int {
	ct := 0
	for head != nil {
		ct++
		head = head.Next
	}
	return ct
}
```