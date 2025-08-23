package _04_maxDepth

// 104. 二叉树的最大深度
// 给定一个二叉树 root ，返回其最大深度。
//
// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：3
// 示例 2：
//
// 输入：root = [1,null,2]
// 输出：2
//
// 提示：
//
// 树中节点的数量在 [0, 104] 区间内。
// -100 <= Node.val <= 100

// 深度优先遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfsTree(root, 1)
}

func dfsTree(t *TreeNode, dep int) int {
	maxDep := dep
	if t.Left != nil {
		leftDep := dfsTree(t.Left, dep+1)
		maxDep = max(maxDep, leftDep)
	}
	if t.Right != nil {
		rightDep := dfsTree(t.Right, dep+1)
		maxDep = max(maxDep, rightDep)
	}
	return maxDep
}

// 广度优先遍历
// 击败: 25.31%
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	maxDep := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		maxDep += 1
	}
	return maxDep
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
