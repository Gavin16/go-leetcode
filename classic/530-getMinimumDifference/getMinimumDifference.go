package _30_getMinimumDifference

import "math"

// 530. 二叉搜索树的最小绝对差
// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 差值是一个正数，其数值等于两值之差的绝对值。
//
// 示例 1：
//
// 输入：root = [4,2,6,1,3]
// 输出：1
// 示例 2：
//
// 输入：root = [1,0,48,null,null,12,49]
// 输出：1
//
// 提示：
//
// 树中节点的数目范围是 [2, 104]
// 0 <= Node.val <= 105

// 二叉搜索树中序遍历转化为 切片 进行判断
// 击败: 83.57%
func getMinimumDifference(root *TreeNode) int {
	slice := make([]int, 0)
	inorderSearch(root, &slice)

	minGap := math.MaxInt32
	for i := 0; i < len(slice)-1; i++ {
		gap := slice[i] - slice[i+1]
		if gap < 0 {
			gap = -gap
		}
		minGap = min(minGap, gap)
	}
	return minGap
}

func inorderSearch(root *TreeNode, slice *[]int) {
	if root == nil {
		return
	}
	inorderSearch(root.Left, slice)
	*slice = append(*slice, root.Val)
	inorderSearch(root.Right, slice)
}

func getMinimumDifference1(root *TreeNode) int {
	slice := make([]int, 0)
	var inorderDFS func(root *TreeNode)
	inorderDFS = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderDFS(root.Left)
		slice = append(slice, root.Val)
		inorderDFS(root.Right)
	}
	inorderDFS(root)

	minGap := math.MaxInt32
	for i := 0; i < len(slice)-1; i++ {
		gap := slice[i] - slice[i+1]
		if gap < 0 {
			gap = -gap
		}
		minGap = min(minGap, gap)
	}
	return minGap
}

// 使用pre 指向前一个元素
func getMinimumDifference2(root *TreeNode) int {
	minGap, pre := math.MaxInt32, -1
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != -1 && root.Val-pre < minGap {
			minGap = root.Val - pre
		}
		pre = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return minGap
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
