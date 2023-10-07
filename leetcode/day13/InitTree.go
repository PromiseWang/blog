package day13

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InsertLeft(val int, root *TreeNode) {
	root.Left = &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
func InsertRight(val int, root *TreeNode) {
	root.Right = &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
