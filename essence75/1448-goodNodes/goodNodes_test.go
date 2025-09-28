package _448_goodNodes

import (
	"testing"
)

func TestGoodNodes(t *testing.T) {
	leftL := &TreeNode{Val: 4}
	leftR := &TreeNode{Val: 2}
	left := &TreeNode{Val: 3, Left: leftL, Right: leftR}
	root := &TreeNode{Val: 3, Left: left, Right: nil}
	root1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	println(goodNodes(root))
	println(goodNodes(root1))
}
