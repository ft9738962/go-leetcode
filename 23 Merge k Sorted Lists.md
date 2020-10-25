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

思路：
1. 将数据全部遍历一边，存放在切片里，建立一个哈希表存放数值和节点的键值对
2. 对切片进行排序
3. 根据切片顺序寻找哈希表中对应的节点连接起来
   
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