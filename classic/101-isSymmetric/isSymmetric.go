package _01_isSymmetric

// 101. 对称二叉树
// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
// 示例 1：
//
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
// 示例 2：
//
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
//
// 提示：
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？

// 递归解法
// 击败: 21.64%
func isSymmetric(root *TreeNode) bool {
	left, right := root.Left, root.Right
	if left == nil && right == nil {
		return true
	}
	if right == nil || left == nil {
		return false
	}
	return dfsCompare(root.Left, root.Right)
}

func dfsCompare(left *TreeNode, right *TreeNode) bool {
	if left.Val != right.Val {
		return false
	}

	sideEqual, midEqual := false, false
	if left.Left != nil && right.Right != nil {
		sideEqual = dfsCompare(left.Left, right.Right)
	} else if left.Left == nil && right.Right == nil {
		sideEqual = true
	}

	if left.Right != nil && right.Left != nil {
		midEqual = dfsCompare(left.Right, right.Left)
	} else if left.Right == nil && right.Left == nil {
		midEqual = true
	}
	return sideEqual && midEqual
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
