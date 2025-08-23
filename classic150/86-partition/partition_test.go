package _6_partition

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	slice := []int{1, 4, 3, 2, 5, 2}
	array := buildListNodeFromArray(slice)
	ans := partition(array, 0)
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
