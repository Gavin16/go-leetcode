package _5_generateTrees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 95-不同的二叉搜索树II
// 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
//
// 示例 1：
//
// 输入：n = 3
// 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
// 示例 2：
//
// 输入：n = 1
// 输出：[[1]]
//
// 提示：
//
// 1 <= n <= 8

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 回溯解法
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return treeMaker(1, n)
}

func treeMaker(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var allTrees []*TreeNode
	for i := start; i <= end; i++ {
		leftTrees := treeMaker(start, i-1)
		rightTrees := treeMaker(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				current := &TreeNode{i, nil, nil}
				current.Left = left
				current.Right = right
				allTrees = append(allTrees, current)
			}
		}
	}
	return allTrees
}
