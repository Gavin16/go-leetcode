package _48_sortList

// 148. 排序链表
// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
// 示例 1：
//
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]
// 示例 2：
//
// 输入：head = [-1,5,3,4,0]
// 输出：[-1,0,3,4,5]
// 示例 3：
//
// 输入：head = []
// 输出：[]
// 提示：
// 链表中节点的数目在范围 [0, 5 * 104] 内
// -105 <= Node.val <= 105
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

// 击败: 14.89%
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	total, curr := 0, head
	for curr != nil {
		total, curr = total+1, curr.Next
	}

	var divide func(mid, n int, head *ListNode) (*ListNode, *ListNode)
	divide = func(ln, n int, head *ListNode) (*ListNode, *ListNode) {
		if head == nil || head.Next == nil {
			return head, nil
		}
		rightHead := head
		leftLast := &ListNode{}
		for i := ln; i > 0; i-- {
			leftLast.Next = rightHead
			rightHead = rightHead.Next
			leftLast = leftLast.Next
		}
		leftLast.Next = nil
		return merge(divide(ln/2, ln, head)), merge(divide((n-ln)/2, n-ln, rightHead))
	}

	left, right := divide(total/2, total, head)
	return merge(left, right)
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	temp := &ListNode{}
	head := temp
	cur1, cur2 := l1, l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			temp.Next = cur1
			temp, cur1 = temp.Next, cur1.Next
		} else {
			temp.Next = cur2
			temp, cur2 = temp.Next, cur2.Next
		}
	}
	for cur1 != nil {
		temp.Next = cur1
		temp, cur1 = temp.Next, cur1.Next
	}
	for cur2 != nil {
		temp.Next = cur2
		temp, cur2 = temp.Next, cur2.Next
	}
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
