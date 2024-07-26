package _37_averageOfLevels

// 637. 二叉树的层平均值
// 给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[3.00000,14.50000,11.00000]
// 解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
// 因此返回 [3, 14.5, 11] 。
// 示例 2:
//
// 输入：root = [3,9,20,15,7]
// 输出：[3.00000,14.50000,11.00000]
//
// 提示：
//
// 树中节点数量在 [1, 104] 范围内
// -231 <= Node.val <= 231 - 1

// 击败: 77.59%
func averageOfLevels(root *TreeNode) []float64 {
	ans := make([]float64, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size, sum := len(queue), 0.0
		for i := 0; i < size; i++ {
			node := queue[i]
			sum += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, sum/float64(size))
		queue = queue[size:]
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
