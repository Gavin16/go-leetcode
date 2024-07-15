package _9_removeNthFromEnd

// 19. 删除链表的倒数第 N 个结点
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
// 示例 2：
//
// 输入：head = [1], n = 1
// 输出：[]
// 示例 3：
//
// 输入：head = [1,2], n = 1
// 输出：[1]
//
// 提示：
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
// 进阶：你能尝试使用一趟扫描实现吗？

// 击败: 21.35%
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	var prev, prePrev *ListNode

	cnt := 1
	for ; fast != nil; cnt += 1 {
		if cnt >= n {
			prePrev, prev = prev, slow
			slow = slow.Next
		}
		fast = fast.Next
	}
	if prePrev == nil {
		return prev.Next
	}
	prePrev.Next = prev.Next
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
