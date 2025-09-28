package _36_lowestCommonAncestor

// 236. 二叉树的最近公共祖先
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
// 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//
// 示例 1：
//
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出：3
// 解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
// 示例 2：
//
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// 输出：5
// 解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
// 示例 3：
//
// 输入：root = [1,2], p = 1, q = 2
// 输出：1
//
// 提示：
//
// 树中节点数目在范围 [2, 105] 内。
// -109 <= Node.val <= 109
// 所有 Node.val 互不相同 。
// p != q
// p 和 q 均存在于给定的二叉树中。

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// 使用广度优先搜索,使用map记录每个节点的父节点
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	parentMap := make(map[int]*TreeNode)
	visit := make(map[int]bool)
	for len(queue) > 0 {
		N := len(queue)
		for i := 0; i < N; i++ {
			node := queue[i]
			left, right := node.Left, node.Right
			if left != nil {
				parentMap[left.Val] = node
				queue = append(queue, left)
			}
			if right != nil {
				parentMap[right.Val] = node
				queue = append(queue, right)
			}
		}
		queue = queue[N:]
	}
	for q != nil {
		visit[q.Val] = true
		q = parentMap[q.Val]
	}
	for p != nil {
		if visit[p.Val] {
			return p
		}
		p = parentMap[p.Val]
	}
	return nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
