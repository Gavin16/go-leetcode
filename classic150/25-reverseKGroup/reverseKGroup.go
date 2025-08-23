package _5_reverseKGroup

// 25. K 个一组翻转链表
// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
// 示例 2：
//
// 输入：head = [1,2,3,4,5], k = 3
// 输出：[3,2,1,4,5]
//
// 提示：
// 链表中的节点数目为 n
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？

// O(1) 额外空间解法
// 击败: 76.69%
func reverseKGroup(head *ListNode, k int) *ListNode {
	curr, ans := head, &ListNode{}

	first, preLast := head, ans
	var last *ListNode

	for cnt := 1; curr != nil; cnt += 1 {
		if cnt%k == 0 {
			last, curr = curr, curr.Next
			// 翻转[first, last] 范围内的节点
			temp, prev := first, first
			for prev != last {
				next := temp.Next
				temp.Next = prev
				prev, temp = temp, next
			}
			first.Next = curr

			// 最后 first -> curr.Next, preLast -> last
			preLast.Next, preLast = last, first
			first = curr
		} else {
			curr = curr.Next
		}
	}
	return ans.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
