package _28_oddEvenList

// 328. 奇偶链表
// 给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别分组，
// 保持它们原有的相对顺序，然后把偶数索引节点分组连接到奇数索引节点分组之后，返回重新排序的链表。
//
// 第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
// 请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
//
// 你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。
// 提示:
//
// n ==  链表中的节点数
// 0 <= n <= 104
// -106 <= Node.val <= 106
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	even := &ListNode{Val: -1, Next: nil}
	evenCurr, oddCurr := even, head
	var oddTail *ListNode

	for oddCurr != nil {
		evenCurr.Next = oddCurr.Next
		evenCurr = evenCurr.Next
		oddTail = oddCurr
		if oddCurr.Next == nil {
			break
		}
		oddCurr.Next = oddCurr.Next.Next
		oddCurr = oddCurr.Next
	}
	oddTail.Next = even.Next
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
