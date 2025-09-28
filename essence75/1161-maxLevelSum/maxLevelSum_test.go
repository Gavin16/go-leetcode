package _161_maxLevelSum

import "testing"

func TestMaxLevelSum(t *testing.T) {
	ll := &TreeNode{Val: 7}
	lr := &TreeNode{Val: -8}
	l := &TreeNode{Val: 7, Left: ll, Right: lr}
	r := &TreeNode{Val: 0}
	root := &TreeNode{Val: 1, Left: l, Right: r}

	println(maxLevelSum(root))
}
