package _3_mergeKLists

// 23. 合并 K 个升序链表
// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//
//	1->4->5,
//	1->3->4,
//	2->6
//
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
// 示例 2：
//
// 输入：lists = []
// 输出：[]
// 示例 3：
//
// 输入：lists = [[]]
// 输出：[]
//
// 提示：
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4

// 分治算法
// 击败: 83.10%
func mergeKLists(lists []*ListNode) *ListNode {

	var mergeTwo func(*ListNode, *ListNode) *ListNode
	mergeTwo = func(l *ListNode, r *ListNode) *ListNode {
		dummyHead := &ListNode{}
		curr := dummyHead
		currL, currR := l, r
		for currL != nil && currR != nil {
			if currL.Val < currR.Val {
				curr.Next = currL
				curr, currL = curr.Next, currL.Next
			} else {
				curr.Next = currR
				curr, currR = curr.Next, currR.Next
			}
		}
		for currL != nil {
			curr.Next = currL
			curr, currL = curr.Next, currL.Next
		}
		for currR != nil {
			curr.Next = currR
			curr, currR = curr.Next, currR.Next
		}
		return dummyHead.Next
	}

	var merge func(list []*ListNode, left, right int) *ListNode
	merge = func(list []*ListNode, left, right int) *ListNode {
		if left == right {
			return nil
		}
		if left+1 == right {
			return list[0]
		}

		mid := len(list) / 2
		l := merge(list[:mid], 0, mid)
		r := merge(list[mid:], 0, len(list)-mid)
		return mergeTwo(l, r)
	}

	return merge(lists, 0, len(lists))
}

type ListNode struct {
	Val  int
	Next *ListNode
}
