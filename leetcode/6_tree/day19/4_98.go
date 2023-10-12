package day19

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
// 有效 二叉搜索树定义如下：
//
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
func isValidBST(root *TreeNode) bool {
	// 错误代码
	//if root.Left == nil && root.Right != nil {
	//	if root.Val < root.Right.Val {
	//		return isValidBST(root.Right)
	//	} else {
	//		return false
	//	}
	//} else if root.Left != nil && root.Right == nil {
	//	if root.Left.Val < root.Val {
	//		return isValidBST(root.Left)
	//	} else {
	//		return false
	//	}
	//} else if root.Left != nil && root.Right != nil {
	//	if root.Left.Val < root.Val && root.Val < root.Right.Val {
	//		return isValidBST(root.Left) && isValidBST(root.Right)
	//	} else {
	//		return false
	//	}
	//} else {
	//	return true
	//}
	var path []int
	inorder(root, &path)
	for i := 1; i < len(path); i++ {
		if path[i-1] >= path[i] {
			return false
		}
	}
	return true
}
func inorder(root *TreeNode, path *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, path)
	*path = append(*path, root.Val)
	inorder(root.Right, path)
}

func T(root *TreeNode) bool {
	return isValidBST(root)
}
