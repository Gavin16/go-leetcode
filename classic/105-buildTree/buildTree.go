package _05_buildTree

// 105. 从前序与中序遍历序列构造二叉树
// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历，
// inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
//
// 示例 1:
//
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
// 示例 2:
//
// 输入: preorder = [-1], inorder = [-1]
// 输出: [-1]
//
// 提示:
//
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder 和 inorder 均 无重复 元素
// inorder 均出现在 preorder
// preorder 保证 为二叉树的前序遍历序列
// inorder 保证 为二叉树的中序遍历序列

// 击败: 42.51%
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
