package _2_deleteDuplicates

// 82. 删除排序链表中的重复元素 II
// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
//
// 示例 1：
//
// 输入：head = [1,2,3,3,4,4,5]
// 输出：[1,2,5]
// 示例 2：
//
// 输入：head = [1,1,1,2,3]
// 输出：[2,3]
//
// 提示：
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序 排列

// 击败 57.37%
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, ans := head, &ListNode{}
	slow, fast, tmp := fast, fast.Next, ans
	for fast != nil {
		if slow.Val == fast.Val {
			for fast != nil && slow.Val == fast.Val {
				fast = fast.Next
			}
			if fast == nil {
				break
			} else if fast.Next == nil {
				tmp.Next, tmp = fast, fast
			}
			slow, fast = fast, fast.Next
		} else {
			tmp.Next = slow
			tmp = tmp.Next
			slow, fast = fast, fast.Next
			if fast == nil {
				tmp.Next, tmp = slow, slow
			}
		}
	}
	tmp.Next = nil
	return ans.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
