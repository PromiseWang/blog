package day21

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var path []int
	getInorderPath(root, &path)
	for {
		index := 0
		indexP := 0
		indexQ := 0
		// 找到当前根结点
		for i := 0; i < len(path); i++ {
			if root.Val == path[i] {
				index = i
			}
			if p.Val == path[i] {
				indexP = i
			}
			if q.Val == path[i] {
				indexQ = i
			}
		}
		if indexP == index {
			return p
		}
		if indexQ == index {
			return q
		}
		// p,q 分散在中序遍历的根结点两侧
		if indexP < index && index < indexQ || indexQ < index && index < indexP {
			return root
		}
		// p,q 在根节点左侧
		if indexP < index && indexQ < index {
			root = root.Left
			continue
		}
		// p,q 在根节点右侧
		if index < indexP && index < indexQ {
			root = root.Right
			continue
		}

	}
}

func getInorderPath(root *TreeNode, path *[]int) {
	if root == nil {
		return
	}
	getInorderPath(root.Left, path)
	*path = append(*path, root.Val)
	getInorderPath(root.Right, path)
}
