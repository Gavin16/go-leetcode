package _095_deleteMiddle

import (
	"testing"
)

func TestDeleteMiddle(t *testing.T) {
	node4 := &ListNode{Val: 5}
	node3 := &ListNode{Val: 4, Next: node4}
	node2 := &ListNode{Val: 3, Next: node3}
	node1 := &ListNode{Val: 2, Next: node2}
	head := &ListNode{Val: 1, Next: node1}

	//head := &ListNode{Val: 1}
	deleteMiddle(head)
}
