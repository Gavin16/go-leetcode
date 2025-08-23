package _2_deleteDuplicates

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	slice1 := []int{1}
	list1 := buildListNodeFromArray(slice1)
	ans1 := deleteDuplicates(list1)
	for ; ans1 != nil; ans1 = ans1.Next {
		fmt.Printf("%v ", ans1.Val)
	}
	fmt.Println()

}

func buildListNodeFromArray(a []int) *ListNode {
	list := &ListNode{}
	curr := list

	for i := 0; i < len(a); i++ {
		curr.Next = &ListNode{Val: a[i]}
		curr = curr.Next
	}
	return list.Next
}
