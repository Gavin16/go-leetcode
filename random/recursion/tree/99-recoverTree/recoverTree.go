package _9_recoverTree

// 99. 恢复二叉搜索树
// 给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
//
// 示例 1：
//
// 输入：root = [1,3,null,null,2]
// 输出：[3,1,null,null,2]
// 解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
// 示例 2：
//
// 输入：root = [3,1,4,null,null,2]
// 输出：[2,1,4,null,null,3]
// 解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
//
// 提示：
//
// 树上节点的数目在范围 [2, 1000] 内
// -231 <= Node.val <= 231 - 1
//
// 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？

// 使用中序遍历左深度优先搜索, 在搜索过程中每个节点的前节点
// 比较前节点和当前节点, 判断是否递增关系 prev.Val < curr.Val
// 若不是则记录首次出现的较大值, 之后每次遇到逆序对，只记录较小值即可
func recoverTree(root *TreeNode) {
	var former, latter, prev *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil && prev.Val > node.Val {
			if former == nil {
				former = prev
			}
			latter = node
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)
	if former != nil && latter != nil {
		former.Val, latter.Val = latter.Val, former.Val
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
