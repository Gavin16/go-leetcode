package _6_partition

// 86. 分隔链表
// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，
// 使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
// 你应当 保留 两个分区中每个节点的初始相对位置。
// 示例 1：
//
// 输入：head = [1,4,3,2,5,2], x = 3
// 输出：[1,2,2,4,3,5]
// 示例 2：
//
// 输入：head = [2,1], x = 2
// 输出：[1,2]
//
// 提示：
//
// 链表中节点的数目在范围 [0, 200] 内
// -100 <= Node.val <= 100
// -200 <= x <= 200

// 击败: 7.21%
func partition(head *ListNode, x int) *ListNode {
	curr, ans := head, &ListNode{}
	tmp := ans

	var back, backHead *ListNode
	for curr != nil {
		if back == nil {
			if curr.Val >= x {
				back, curr = curr, curr.Next
				backHead = back
			} else {
				tmp.Next, tmp, curr = curr, curr, curr.Next
			}
		} else {
			if curr.Val >= x {
				back.Next, curr = curr, curr.Next
				back = back.Next
			} else {
				tmp.Next, curr = curr, curr.Next
				tmp, back.Next = tmp.Next, curr
			}
		}
	}
	tmp.Next = backHead
	return ans.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
