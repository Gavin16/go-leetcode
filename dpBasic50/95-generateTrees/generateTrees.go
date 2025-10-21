package _5_generateTrees

// 95. 不同的二叉搜索树 II
// 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。
// 可以按 任意顺序 返回答案。
//
// 示例 1：
// 输入：n = 3
// 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
//
// 示例 2：
// 输入：n = 1
// 输出：[[1]]
//
// 提示：
// 1 <= n <= 8

func generateTrees(n int) []*TreeNode {
	return dfsGenTrees(1, n)
}

func dfsGenTrees(s, t int) []*TreeNode {
	if s > t {
		return []*TreeNode{nil}
	}
	var roots []*TreeNode
	for i := s; i <= t; i++ {
		lefts := dfsGenTrees(s, i-1)
		rights := dfsGenTrees(i+1, t)
		for _, l := range lefts {
			for _, r := range rights {
				root := &TreeNode{Val: i}
				root.Left = l
				root.Right = r
				roots = append(roots, root)
			}
		}
	}
	return roots
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
