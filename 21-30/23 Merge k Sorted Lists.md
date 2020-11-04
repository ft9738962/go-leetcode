You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

 

**Example 1:**
```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
```
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

**Example 2:**
```
Input: lists = []
Output: []
```
**Example 3:**
```
Input: lists = [[]]
Output: []
```

**Constraints:**

- k == lists.length
- 0 <= k <= 10^4
- 0 <= lists[i].length <= 500
- -10^4 <= lists[i][j] <= 10^4
- lists[i] is sorted in ascending order.
- The sum of lists[i].length won't exceed 10^4.

### 思路1：
1. 将数据全部遍历一边，存放在切片里，建立一个哈希表存放数值和节点的键值对
2. 对切片进行排序
3. 根据切片顺序寻找哈希表中对应的节点连接起来

优点：不如思路2
缺点：需要额外的哈希表存放键值对关系
   
```go
// runtime 12ms 70.4% memory 6.4MB 10.4%
func mergeKLists(lists []*ListNode) *ListNode {

	var emp *ListNode
	if len(lists) == 0 {
		return emp
	}

	vals, valMap, nonEmptyListCount := firstOrderLists(lists, emp)
	if nonEmptyListCount == 0 {
		return emp
	}

	sort.Ints(vals)
	var resultHead, curResult *ListNode

	for i := 0; i < len(vals); i++ {
		if i > 0 {
			if vals[i] == vals[i-1] {
				continue
			}
		}

		nodes := valMap[vals[i]]
		for j := 0; j < len(nodes); j++ {
			if resultHead == emp {
				resultHead = nodes[j]
				curResult = resultHead
			} else {
				curResult.Next = nodes[j]
				curResult = curResult.Next
			}
		}
	}

	return resultHead
}

func firstOrderLists(lists []*ListNode, emp *ListNode) ([]int, map[int][]*ListNode, int) {
	vals := []int{}
	valMap := map[int][]*ListNode{}
	nonEmptyListCount := 0
	var list *ListNode
	for i := 0; i < len(lists); i++ {
		if lists[i] != emp {
			list = lists[i]
			for list.Next != nil {
				vals = append(vals, list.Val)
				if _, ok := valMap[list.Val]; ok != false {
					valMap[list.Val] = append(valMap[list.Val], list)
				} else {
					valMap[list.Val] = []*ListNode{list}
				}
				list = list.Next

			}
			vals = append(vals, list.Val)
			if _, ok := valMap[list.Val]; ok != false {
				valMap[list.Val] = append(valMap[list.Val], list)
			} else {
				valMap[list.Val] = []*ListNode{list}
			}
			nonEmptyListCount++
		}
	}
	return vals, valMap, nonEmptyListCount
}
```

### 思路2：
1. 每次取2条链表进行合并，将合并好的新列表插入原切片，抛弃原来的2条链表
2. 重复过程1，直到只剩1条链表
3. 返回链表头
4. 算法复杂度O(nlogk) k为链表的数量

优点：较思路1节约了哈希表

```go
// runtime 8ms 95.65% memory 5.7MB 10.87% 
func mergeKLists(lists []*ListNode) *ListNode {
	var emp *ListNode

	if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 0 {
		return emp
	}

	if len(lists) > 2 {
		for len(lists) > 2 {
			lists = append(lists, mergeTwoLists(lists[0], lists[1]))[2:]
		}
	}
	return mergeTwoLists(lists[0], lists[1])
}

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	prehead := &ListNode{Val: 0}
	cur := prehead

	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next

	}

	if a == nil {
		cur.Next = b
	} else {
		cur.Next = a
	}
	return prehead.Next
}

func printVals(head *ListNode) {
	fmt.Println("start")
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("end")
}
```