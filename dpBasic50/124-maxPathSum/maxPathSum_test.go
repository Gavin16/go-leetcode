package _24_maxPathSum

import "testing"

func TestMaxPathSum(t *testing.T) {
	rr := &TreeNode{Val: 2,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: -6},
			Left:  &TreeNode{Val: -6, Left: &TreeNode{Val: -6}},
		},
	}
	right := &TreeNode{Val: -3, Left: &TreeNode{Val: -6}, Right: rr}
	root := &TreeNode{Val: 9, Left: &TreeNode{Val: 6}, Right: right}
	println(maxPathSum(root))
}
