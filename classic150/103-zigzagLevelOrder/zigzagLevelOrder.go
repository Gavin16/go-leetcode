package _03_zigzagLevelOrder

// 103. 二叉树的锯齿形层序遍历
// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
// 示例 2：
//
// 输入：root = [1]
// 输出：[[1]]
// 示例 3：
//
// 输入：root = []
// 输出：[]
//
// 提示：
//
// 树中节点数目在范围 [0, 2000] 内
// -100 <= Node.val <= 100

// 击败: 11.44%
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result, level := make([][]int, 0), 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size, slice := len(queue), make([]int, 0)
		level += 1
		for i := 0; i < size; i++ {
			node := queue[i]
			slice = append(slice, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		if level%2 == 0 {
			reverse(slice)
		}
		result = append(result, slice)
	}
	return result
}

func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
