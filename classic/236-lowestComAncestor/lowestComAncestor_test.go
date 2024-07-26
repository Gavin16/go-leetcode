package _36_lowestComAncestor

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	preorder := []int{3, 5, 6, 2, 7, 4, 1, 0, 8}
	inorder := []int{6, 5, 7, 2, 4, 3, 0, 1, 8}
	tree := buildTree(preorder, inorder)

	ans := lowestCommonAncestor(tree, tree.Left, tree.Right)
	fmt.Printf("%v\n", ans)
}

// 先序遍历 + 中序遍历构造 二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	return dfsBuildTree(preorder, inorder, 0)
}

func dfsBuildTree(preorder []int, inorder []int, preId int) *TreeNode {
	node := &TreeNode{Val: preorder[preId]}

	pos := findInorderPos(inorder, preorder[preId])
	leftI, lenLeft := inorder[:pos], len(inorder[:pos])
	leftP := preorder[preId+1 : preId+1+lenLeft]
	if len(leftI) > 0 {
		leftNode := dfsBuildTree(leftP, leftI, 0)
		node.Left = leftNode
	}

	rightI := inorder[pos+1:]
	rightP := preorder[preId+1+lenLeft:]
	if len(rightI) > 0 {
		rightNode := dfsBuildTree(rightP, rightI, 0)
		node.Right = rightNode
	}
	return node
}

func findInorderPos(inorder []int, val int) int {
	for j := 0; j < len(inorder); j++ {
		if inorder[j] == val {
			return j
		}
	}
	return -1
}
