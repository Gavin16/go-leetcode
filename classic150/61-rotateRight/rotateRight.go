package _1_rotateRight

// 61. 旋转链表
// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[4,5,1,2,3]
// 示例 2：
//
// 输入：head = [0,1,2], k = 4
// 输出：[2,0,1]
//
// 提示：
//
// 链表中节点的数目在范围 [0, 500] 内
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 109

// 击败: 15.14%
func rotateRight(head *ListNode, k int) *ListNode {
	cursor, cnt := head, 0
	for cursor != nil {
		cursor, cnt = cursor.Next, cnt+1
	}
	if cnt == 0 || k%cnt == 0 {
		return head
	}
	offset := k % cnt
	cursor = head
	for i := 1; i < cnt-offset; i++ {
		cursor = cursor.Next
	}
	newHead, last := cursor.Next, cursor.Next
	cursor.Next = nil
	for last.Next != nil {
		last = last.Next
	}
	last.Next = head

	return newHead

}

type ListNode struct {
	Val  int
	Next *ListNode
}
