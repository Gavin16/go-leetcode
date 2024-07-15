package _2_reverseBetween

// 92. 反转链表 II
// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
// 请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]
// 示例 2：
//
// 输入：head = [5], left = 1, right = 1
// 输出：[5]
//
// 提示：
//
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
// 进阶： 你可以使用一趟扫描完成反转吗？

// 使用一趟扫描完成反转
// 击败: 24.69%
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	curr := head
	var prev, post, preFirst *ListNode
	var first, last *ListNode
	for cnt := 1; cnt <= right; cnt += 1 {
		if cnt < left || cnt > right {
			prev, curr = curr, curr.Next
		} else {
			if cnt == left {
				first, preFirst = curr, prev
			}
			if cnt == right {
				last, post = curr, curr.Next
			}
			next := curr.Next
			curr.Next, prev = prev, curr
			curr = next
		}
	}
	first.Next = post
	if preFirst != nil {
		preFirst.Next = last
		return head
	}
	return last
}

type ListNode struct {
	Val  int
	Next *ListNode
}
