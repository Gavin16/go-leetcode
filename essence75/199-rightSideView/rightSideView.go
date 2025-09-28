package _99_rightSideView

// 199. 二叉树的右视图
// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
// 示例 1：
// 输入：root = [1,2,3,null,5,null,4]
// 输出：[1,3,4]
// 解释：
//
// 示例 2：
// 输入：root = [1,2,3,4,null,null,null,5]
// 输出：[1,3,4,5]
// 解释：
//
// 示例 3：
// 输入：root = [1,null,3]
// 输出：[1,3]
//
// 示例 4：
// 输入：root = []
// 输出：[]
//
// 提示:
// 二叉树的节点个数的范围是 [0,100]
// -100 <= Node.val <= 100

func rightSideView(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	if root == nil {
		return []int{}
	}
	var ans []int
	queue = append(queue, root)
	for len(queue) > 0 {
		N := len(queue)
		for i := 0; i < N; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, queue[N-1].Val)
		queue = queue[N:]
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
