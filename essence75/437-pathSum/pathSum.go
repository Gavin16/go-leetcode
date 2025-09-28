package _37_pathSum

// 437. 路径总和 III
// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，
// 求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
// 提示:
//
// 二叉树的节点个数的范围是 [0,1000]
// -109 <= Node.val <= 109
// -1000 <= targetSum <= 1000

func pathSum(root *TreeNode, targetSum int) int {
	slice := make([]*TreeNode, 0)
	sumCnt := 0
	var dfs func(node *TreeNode, path []*TreeNode)
	dfs = func(node *TreeNode, path []*TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node)
		sum := 0
		for i := len(path) - 1; i >= 0; i-- {
			sum += path[i].Val
			if sum == targetSum {
				sumCnt += 1
			}
		}
		dfs(node.Left, path)
		dfs(node.Right, path)
		path = path[:len(path)-1]
	}
	dfs(root, slice)
	return sumCnt
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
