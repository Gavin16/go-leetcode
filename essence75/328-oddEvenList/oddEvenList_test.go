package _28_oddEvenList

import "testing"

func TestOddEvenList(t *testing.T) {
	//node4 := &ListNode{Val: 5, Next: nil}
	//node3 := &ListNode{Val: 4, Next: node4}
	//node2 := &ListNode{Val: 3, Next: node3}
	node1 := &ListNode{Val: 2, Next: nil}
	header := &ListNode{Val: 1, Next: node1}

	oddEvenList(header)
}
