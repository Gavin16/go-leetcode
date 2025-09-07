package _35_lowestCommonAncestor

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	tree := &TreeNode{Val: 6,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	p, q := tree.Left, tree.Right

	//p, q := tree.Left.Left, tree.Left.Right.Left
	//if ancestor := lowestCommonAncestor(tree, p, q); ancestor != nil {
	//	println(ancestor.Val)
	//} else {
	//	println("未找到最近祖先")
	//}

	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	p, q = root, root.Left
	if ancestor := lowestCommonAncestor(root, p, q); ancestor != nil {
		println(ancestor.Val)
	} else {
		println("未找到最近祖先")
	}

}
