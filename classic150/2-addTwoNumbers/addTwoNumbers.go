package __addTwoNumbers

// 2. 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例 1：
//
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 示例 2：
//
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
// 示例 3：
//
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]
//
// 提示：
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零

// 击败: 47.66%
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel, carry := &ListNode{Val: -1}, 0
	ans := sentinel
	for l1 != nil && l2 != nil {
		ans.Next = &ListNode{Val: (carry + l1.Val + l2.Val) % 10}
		carry, ans = (carry+l1.Val+l2.Val)/10, ans.Next
		l1, l2 = l1.Next, l2.Next
	}

	for l1 != nil {
		ans.Next = &ListNode{Val: (carry + l1.Val) % 10}
		l1, carry, ans = l1.Next, (carry+l1.Val)/10, ans.Next
	}
	for l2 != nil {
		ans.Next = &ListNode{Val: (carry + l2.Val) % 10}
		l2, carry, ans = l2.Next, (carry+l2.Val)/10, ans.Next
	}
	if carry > 0 {
		ans.Next = &ListNode{Val: carry}
	}
	return sentinel.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
