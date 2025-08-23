package _48_sortList

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	head1 := []int{4, 2, 1, 3}
	l1 := convertSliceToListNodes(head1)
	list := sortList(l1)
	for list != nil {
		fmt.Printf("%v ", list.Val)
		list = list.Next
	}
	fmt.Println()

	head2 := []int{-1, 5, 3, 4, 0}
	l2 := convertSliceToListNodes(head2)
	list1 := sortList(l2)
	for list1 != nil {
		fmt.Printf("%v ", list1.Val)
		list1 = list1.Next
	}
	fmt.Println()

}

func convertSliceToListNodes(s []int) *ListNode {
	head := &ListNode{}
	curr := head
	for i := 0; i < len(s); i++ {
		curr.Next = &ListNode{Val: s[i]}
		curr = curr.Next
	}
	return head.Next
}
