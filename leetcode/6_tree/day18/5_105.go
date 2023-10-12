package day18

// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
func buildTree2(preorder []int, inorder []int) *TreeNode {
	// 1.前序数组为0, 空节点
	if len(preorder) == 0 {
		return nil
	}
	// 2.前序数组第一个元素为节点元素
	rootValue := preorder[0]
	root := &TreeNode{
		Val:   rootValue,
		Left:  nil,
		Right: nil,
	}
	if len(preorder) == 1 {
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
	// 5.切前序数组
	leftPreorder := preorder[1 : len(leftInorder)+1]
	rightPreorder := preorder[len(leftPreorder)+1:]
	// 6.递归处理左区间 右区间
	root.Left = buildTree2(leftPreorder, leftInorder)
	root.Right = buildTree2(rightPreorder, rightInorder)
	return root
}
