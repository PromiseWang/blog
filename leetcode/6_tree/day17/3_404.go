package day17

// 给定二叉树的根节点 root ，返回所有左叶子之和。
func sumOfLeftLeaves(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}
	sum := 0
	getSum(root, &sum, false)
	return sum
}

func getSum(root *TreeNode, sum *int, flag bool) { // flag为true, 代表左节点; 否则是右结点
	if root.Left == nil && root.Right == nil && flag {
		*sum += root.Val
		return
	}
	if root.Left != nil {
		getSum(root.Left, sum, true)
	}
	if root.Right != nil {
		getSum(root.Right, sum, false)
	}
}

func T(root *TreeNode) int {
	return sumOfLeftLeaves(root)
}
