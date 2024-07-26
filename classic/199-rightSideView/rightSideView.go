package _99_rightSideView

// 199. 二叉树的右视图
// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
// 示例 1:
//
// 输入: [1,2,3,null,5,null,4]
// 输出: [1,3,4]
// 示例 2:
//
// 输入: [1,null,3]
// 输出: [1,3]
// 示例 3:
//
// 输入: []
// 输出: []
//
// 提示:
//
// 二叉树的节点个数的范围是 [0,100]
// -100 <= Node.val <= 100

// 广度优先搜索每次取最后一个元素
// 击败: 100.00%
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue, result := make([]*TreeNode, 0), make([]int, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if i == size-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
