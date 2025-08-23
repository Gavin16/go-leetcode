package _24_maxPathSum

// 124. 二叉树中的最大路径和
// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。
// 同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//
// 路径和 是路径中各节点值的总和。
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
//
// 示例 1：
//
// 输入：root = [1,2,3]
// 输出：6
// 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
// 示例 2：
//
// 输入：root = [-10,9,20,null,null,15,7]
// 输出：42
// 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
//
// 提示：
//
// 树中节点数目范围是 [1, 3 * 104]
// -1000 <= Node.val <= 1000

// 回溯算法
// 对于叶子节点 仅需与maxSum初始值(-1001)比较大小
// 对于非叶子节点，将左边最大值 和 右边最大值 加上自身值 与最大值maxSum比较
//
//	同时非叶子节点返回上一级时, 返回 max(val+maxLeft, val+maxRight)
//
// 完成深度优先遍历后, maxSum得到的即为 最大路径和
// 击败: 25.01%
func maxPathSum(root *TreeNode) int {
	maxSum := -1001
	dfsTreePath(root, &maxSum)
	return maxSum
}

func dfsTreePath(root *TreeNode, maxSum *int) int {
	if root.Left == nil && root.Right == nil {
		*maxSum = max(*maxSum, root.Val)
		return root.Val
	}
	leftSum, rightSum := 0, 0
	if root.Left != nil {
		leftSum = dfsTreePath(root.Left, maxSum)
	}
	if root.Right != nil {
		rightSum = dfsTreePath(root.Right, maxSum)
	}
	currMax := root.Val
	if leftSum > 0 {
		currMax += leftSum
	}
	if rightSum > 0 {
		currMax += rightSum
	}
	*maxSum = max(*maxSum, currMax)
	return max(root.Val, max(root.Val+leftSum, root.Val+rightSum))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
