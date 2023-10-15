package day22

// 给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。
// 修剪树 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。
//
// 所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	path := []int{}
	pre(root, &path)
	for _, v := range path {
		if !(low <= v && v <= high) {
			root = deleteNode(root, v)
		}
	}
	return root
}

func pre(root *TreeNode, path *[]int) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	pre(root.Left, path)
	pre(root.Right, path)
}

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
