package day17

// 平衡二叉树
func isBalanced(root *TreeNode) bool {
	return backTracking(root) != -1
}
func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func backTracking(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// left child
	leftHeight := backTracking(root.Left)
	if leftHeight == -1 {
		return -1
	}

	// right child
	rightHeight := backTracking(root.Right)
	if rightHeight == -1 {
		return -1
	}

	if rightHeight-leftHeight > 1 || rightHeight-leftHeight < -1 {
		return -1
	} else {
		return maxNum(rightHeight, leftHeight) + 1
	}
}
