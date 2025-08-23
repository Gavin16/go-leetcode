package _06_buildTree

// 106. 从中序与后序遍历序列构造二叉树
// 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
//
// 示例 1:
//
// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 输出：[3,9,20,null,null,15,7]
// 示例 2:
//
// 输入：inorder = [-1], postorder = [-1]
// 输出：[-1]
//
// 提示:
//
// 1 <= inorder.length <= 3000
// postorder.length == inorder.length
// -3000 <= inorder[i], postorder[i] <= 3000
// inorder 和 postorder 都由 不同 的值组成
// postorder 中每一个值都在 inorder 中
// inorder 保证是树的中序遍历
// postorder 保证是树的后序遍历

// 击败: 26.39%
func buildTree(inorder []int, postorder []int) *TreeNode {
	seqLen := len(inorder)
	return dfsBuildTree(inorder, postorder, seqLen-1)
}

func dfsBuildTree(inorder []int, postorder []int, postId int) *TreeNode {
	value := postorder[postId]
	node := &TreeNode{Val: value}
	iPos := findInorderPos(inorder, value)
	leftInorder, rightInorder := inorder[0:iPos], inorder[iPos+1:]
	leftPostOrd := postorder[0:len(leftInorder)]
	rightPostOrd := postorder[len(leftInorder) : len(postorder)-1]

	if len(leftInorder) > 0 {
		left := dfsBuildTree(leftInorder, leftPostOrd, len(leftPostOrd)-1)
		node.Left = left
	}
	if len(rightInorder) > 0 {
		right := dfsBuildTree(rightInorder, rightPostOrd, len(rightPostOrd)-1)
		node.Right = right
	}
	return node
}

// 没找到则返回 -1
func findInorderPos(inorder []int, val int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == val {
			return i
		}
	}
	return -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
