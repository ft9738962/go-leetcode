## 2. Add Two Numbers ##

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```


先判断有没有链只有一个节点0，有的话直接返回另一条链
后面直接在l1链上操作： 
1. 当两条链都有下一个节点时：把l2的数加过来，并记录是否有进位带到下一个节点
2. 当两条链同时没有下一个节点时，把l2的数加过来，看是否有进位，有进位则再l1后面加一个值为1的点，返回l1的头部节点
3. 当有一条链没有下一个节点时，把目前节点的计算处理完，将进位带过去，进入只处理另一条链的后续
4. 把处理好的后续节点拼接到两条链同时处理的最后节点上，返回l1链
```go
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	} else if l2.Val == 0 && l2.Next == nil {
		return l1
	}

	var carry int //v1,v2保存链表值，carry保存进位符
	cur := l1

	for (cur.Next != nil) && (l2.Next != nil) {
		cur.Val, carry = addNodeVal(cur, l2, carry)
		cur, l2 = cur.Next, l2.Next
	}

	cur.Val, carry = addNodeVal(cur, l2, carry)
	if cur.Next == nil && l2.Next == nil {
		if carry == 1 {
			cur.Next = &ListNode{Val: 1}
		}
	} else if cur.Next == nil {
		cur.Next = singleLinkedAdd(l2.Next, carry)
	} else if l2.Next == nil {
		cur.Next = singleLinkedAdd(cur.Next, carry)
	}
	return l1
}
```

```go
func addNodeVal(cur *ListNode, other *ListNode, carry int) (int, int) {
	return (cur.Val + other.Val + carry) % 10, (cur.Val + other.Val + carry) / 10
}

func singleLinkedAdd(l *ListNode, carry int) *ListNode {
	cur := l
	for {
		cur.Val, carry = addNodeVal(&ListNode{Val: 0}, cur, carry)
		if carry == 1 && cur.Next == nil {
			cur.Next = &ListNode{Val: 1, Next: nil}
			return l
		}
		if cur.Next == nil {
			cur.Val += carry
			return l
		}
		cur = cur.Next
	}
}
```