package day18

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Init(root *TreeNode) *TreeNode {
	root = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left.Left = &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right = &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	root.Right.Right = &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	return root
}
