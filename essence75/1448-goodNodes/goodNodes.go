package _448_goodNodes

import "math"

// 1448. 统计二叉树中好节点的数目
// 给你一棵根为 root 的二叉树，请你返回二叉树中好节点的数目。
//
// 「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值。
// 提示：
//
// 二叉树中节点数目范围是 [1, 10^5]
// 每个节点权值的范围是 [-10^4, 10^4]

func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, m int) int
	dfs = func(node *TreeNode, m int) int {
		if node == nil {
			return 0
		}
		ret := 0
		if m <= node.Val {
			m = node.Val
			ret = 1
		}
		left := dfs(node.Left, m)
		right := dfs(node.Right, m)
		return ret + left + right
	}
	return dfs(root, math.MinInt)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
