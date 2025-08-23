package _36_lowestComAncestor

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

// 回溯算法
// 递归返回时判断
// 击败: 64.96%
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	result := make([]*TreeNode, 0)
	dfs(root, p, q, &result)
	return result[0]
}

func dfs(root, p, q *TreeNode, slice *[]*TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		if root == p || root == q {
			return true
		}
		return false
	}
	leftExists, rightExists := false, false
	currEqual := root == p || root == q

	if root.Left != nil {
		leftExists = dfs(root.Left, p, q, slice)
	}
	if root.Right != nil {
		rightExists = dfs(root.Right, p, q, slice)
	}

	if (currEqual && (leftExists || rightExists)) || (leftExists && rightExists) {
		if len(*slice) == 0 {
			*slice = append(*slice, root)
		}
		return true
	} else if currEqual || leftExists || rightExists {
		return true
	}
	return false
}

// 解法2: 直接搜索
// 若 p,q 是直接父子节点的关系, 返回父节点就好，不需要一直走到叶子节点
// 若 p,q 在不同的路径上, 祖先节点左右节点是否为 nil 都为nil返回nil, 有一个不为nil 返回它 都不为nil 则返回该祖先节点
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
