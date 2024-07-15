package _1_mergeTwoLists

// 21. 合并两个有序链表
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例 1：
//
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// 示例 2：
//
// 输入：l1 = [], l2 = []
// 输出：[]
// 示例 3：
//
// 输入：l1 = [], l2 = [0]
// 输出：[0]
//
// 提示：
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列

// 击败: 11.76%
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	sentinel := &ListNode{Val: -101}
	curr := sentinel
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			curr.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}
	return sentinel.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
