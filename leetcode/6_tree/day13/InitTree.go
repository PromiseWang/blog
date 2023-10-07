package day13

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Operator interface {
	InsertLeft()
}
