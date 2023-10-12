package day19

// 给定二叉搜索树（BST）的根节点 root 和一个整数值 val。
//
// 你需要在 BST 中找到节点值等于 val 的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 null 。
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	var node *TreeNode
	if root.Val == val {
		node = root
		return node
	} else if val < root.Val {
		node = searchBST(root.Left, val)
	} else {
		node = searchBST(root.Right, val)
	}
	return node
}
