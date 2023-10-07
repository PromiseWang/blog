package day14

func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right == nil {
		return false
	}
	if left == nil && right != nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return compare(left.Left, right.Right) && compare(left.Right, right.Left)

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func T(root *TreeNode) bool {
	return isSymmetric(root)
}
