package day21

// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
//
// 一般来说，删除节点可分为两个步骤：
//
// 首先找到需要删除的节点；
// 如果找到了，删除它。
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 1. 没找到删除节点
	if root == nil {
		return root
	}
	if root.Val == key {
		// 2. 该节点是叶子节点: 直接删除
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 3. 左空右不空: 右节点补位
		if root.Left == nil && root.Right != nil {
			node := root.Right
			return node
		}
		// 4. 左不空右空: 左节点补位
		if root.Left != nil && root.Right == nil {
			node := root.Left
			return node
		}
		// 5. 都不空: 则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点。
		if root.Left != nil && root.Right != nil {
			node := root.Right
			for node.Left != nil {
				node = node.Left
			}
			node.Left = root.Left
			root = root.Right
			return root
		}
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func T(root *TreeNode, key int) *TreeNode {
	return deleteNode(root, key)
}
