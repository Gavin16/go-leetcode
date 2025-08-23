package _2_reverseBetween

import (
	"fmt"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	nodes1 := buildListNodeFromArray(arr1)
	between1 := reverseBetween(nodes1, 1, 5)
	for ; between1 != nil; between1 = between1.Next {
		fmt.Printf("%v\n", between1.Val)
	}

	arr2 := []int{3}
	nodes2 := buildListNodeFromArray(arr2)
	between2 := reverseBetween(nodes2, 1, 1)
	for ; between2 != nil; between2 = between2.Next {
		fmt.Printf("%v\n", between2.Val)
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
