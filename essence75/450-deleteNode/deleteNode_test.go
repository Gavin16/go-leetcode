package _50_deleteNode

import "testing"

func TestDeleteNode(t *testing.T) {
	lr := &TreeNode{Val: 40}
	l := &TreeNode{Val: 30, Right: lr}
	rl := &TreeNode{Val: 60}
	rr := &TreeNode{Val: 80}
	r := &TreeNode{Val: 70, Left: rl, Right: rr}
	root := &TreeNode{Val: 50, Left: l, Right: r}
	ans := deleteNode(root, 50)
	println(ans)
}
