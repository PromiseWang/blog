package day18

// 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 1.后序数组长度为0, 空节点
	if len(postorder) == 0 {
		return nil
	}
	// 2.后序数组最后一个元素为父节点元素
	rootValue := postorder[len(postorder)-1]
	root := &TreeNode{
		Val:   rootValue,
		Left:  nil,
		Right: nil,
	}
	if len(postorder) == 1 {
		return root
	}
	// 3.寻找中序数组位置  作切割点
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == rootValue {
			break
		}
	}

	// 4.切中序数组
	leftInorder := inorder[:index]
	rightInorder := inorder[index+1:]
	// 5.切后序数组
	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftPostorder) : len(leftPostorder)+len(rightInorder)]
	// 6.递归处理左区间 右区间
	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}
