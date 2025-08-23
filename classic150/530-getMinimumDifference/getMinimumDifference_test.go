package _30_getMinimumDifference

import "testing"

func TestGetMinimumDifference(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}

	println(getMinimumDifference(root))

}
