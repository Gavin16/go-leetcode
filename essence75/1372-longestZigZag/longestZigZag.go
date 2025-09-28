package _372_longestZigZag

// 1372. 二叉树中的最长交错路径
// 给你一棵以 root 为根的二叉树，二叉树中的交错路径定义如下：
//
// 选择二叉树中 任意 节点和一个方向（左或者右）。
// 如果前进方向为右，那么移动到当前节点的的右子节点，否则移动到它的左子节点。
// 改变前进方向：左变右或者右变左。
// 重复第二步和第三步，直到你在树中无法继续移动。
// 交错路径的长度定义为：访问过的节点数目 - 1（单个节点的路径长度为 0 ）。
//
// 请你返回给定树中最长 交错路径 的长度。
// 提示：
//
// 每棵树最多有 50000 个节点。
// 每个节点的值在 [1, 100] 之间。

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLen := 0
	var dfs func(node *TreeNode, isLeft bool, depth int)
	dfs = func(node *TreeNode, isLeft bool, depth int) {
		maxLen = max(maxLen, depth)
		if isLeft {
			if node.Left != nil {
				dfs(node.Left, true, 1)
			}
			if node.Right != nil {
				dfs(node.Right, false, depth+1)
			}
		} else {
			if node.Right != nil {
				dfs(node.Right, false, 1)
			}
			if node.Left != nil {
				dfs(node.Left, true, depth+1)
			}
		}
	}
	dfs(root, true, 0)
	dfs(root, false, 0)
	return maxLen
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
