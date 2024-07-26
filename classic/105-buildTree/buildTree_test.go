package _05_buildTree

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {

	preorder, inorder := []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)
	printTree(tree)

	preorder1, inorder1 := []int{-1}, []int{-1}
	tree1 := buildTree(preorder1, inorder1)
	printTree(tree1)
}

func printTree(tree *TreeNode) {
	if tree == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, tree)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			t := queue[i]
			if t == nil {
				fmt.Printf("%v  ", "nil")
			} else {
				fmt.Printf("%v  ", t.Val)
				queue = append(queue, t.Left, t.Right)
			}
		}
		queue = queue[size:]
		fmt.Printf("\n")
	}
}
