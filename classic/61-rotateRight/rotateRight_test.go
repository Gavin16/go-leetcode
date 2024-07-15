package _1_rotateRight

import (
	"fmt"
	"testing"
)

func TestRotateRight(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	list := buildListNodeFromArray(slice)
	ans := rotateRight(list, 1)
	for ; ans != nil; ans = ans.Next {
		fmt.Printf("%v ", ans.Val)
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
