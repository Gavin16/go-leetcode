package _00_isSameTree

// 100. 相同的树
// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
// 示例 1：
//
// 输入：p = [1,2,3], q = [1,2,3]
// 输出：true
// 示例 2：
//
// 输入：p = [1,2], q = [1,null,2]
// 输出：false
// 示例 3：
//
// 输入：p = [1,2,1], q = [1,1,2]
// 输出：false
//
// 提示：
//
// 两棵树上的节点数目都在范围 [0, 100] 内
// -104 <= Node.val <= 104

// 广度优先搜索判断
// 击败:  100.00%
func isSameTree(p *TreeNode, q *TreeNode) bool {
	pQueue, qQueue := make([]*TreeNode, 0), make([]*TreeNode, 0)
	if p == nil && q == nil {
		return true
	} else if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}
	pQueue, qQueue = append(pQueue, p), append(qQueue, q)
	for len(pQueue) > 0 && len(qQueue) > 0 {
		pSize, qSize := len(pQueue), len(qQueue)
		if pSize != qSize {
			return false
		}
		for i := 0; i < pSize; i++ {
			curP, curQ := pQueue[0], qQueue[0]
			if curP.Val != curQ.Val {
				return false
			}
			pQueue, qQueue = pQueue[1:], qQueue[1:]
			pLeft, qLeft := curP.Left, curQ.Left
			if (pLeft == nil && qLeft == nil) || (pLeft != nil && qLeft != nil) {
				if pLeft != nil {
					pQueue = append(pQueue, pLeft)
					qQueue = append(qQueue, qLeft)
				}
			} else {
				return false
			}
			pRight, qRight := curP.Right, curQ.Right
			if (pRight == nil && qRight == nil) || (pRight != nil && qRight != nil) {
				if pRight != nil {
					pQueue = append(pQueue, pRight)
					qQueue = append(qQueue, qRight)
				}
			} else {
				return false
			}
		}
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
