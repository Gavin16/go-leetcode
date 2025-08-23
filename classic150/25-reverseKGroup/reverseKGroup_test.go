package _5_reverseKGroup

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	list1 := buildListNodeFromArray(arr1)
	ans1 := reverseKGroup(list1, 5)
	for ; ans1 != nil; ans1 = ans1.Next {
		fmt.Printf("%v\n", ans1.Val)
	}

	arr2 := []int{5}
	list2 := buildListNodeFromArray(arr2)
	ans2 := reverseKGroup(list2, 1)
	for ; ans2 != nil; ans2 = ans2.Next {
		fmt.Printf("%v\n", ans2.Val)
	}
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
