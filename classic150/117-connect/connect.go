package _17_connect

// 117. 填充每个节点的下一个右侧节点指针 II
// 给定一个二叉树：
//
//	struct Node {
//	 int val;
//	 Node *left;
//	 Node *right;
//	 Node *next;
//	}
//
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
//
// 初始状态下，所有 next 指针都被设置为 NULL 。
//
// 示例 1：
//
// 输入：root = [1,2,3,4,5,null,7]
// 输出：[1,#,2,3,#,4,5,7,#]
// 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化输出按层序遍历顺序（由 next 指针连接），'#' 表示每层的末尾。
// 示例 2：
//
// 输入：root = []
// 输出：[]
//
// 提示：
//
// 树中的节点数在范围 [0, 6000] 内
// -100 <= Node.val <= 100
//
// 进阶：
// 你只能使用常量级额外空间。
// 使用递归解题也符合要求，本题中递归程序的隐式栈空间不计入额外空间复杂度。

// 使用广度优先搜索, 空间复杂度 O(n)
// 击败: 67.19%
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if i < size-1 {
				node.Next = queue[i+1]
			} else {
				node.Next = nil
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
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
