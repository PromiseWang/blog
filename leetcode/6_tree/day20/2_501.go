package day20

// 给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
//
// 如果树中有不止一个众数，可以按 任意顺序 返回。
//
// 假定 BST 满足如下定义：
//
// 结点左子树中所含节点的值 小于等于 当前节点的值
// 结点右子树中所含节点的值 大于等于 当前节点的值
// 左子树和右子树都是二叉搜索树
func findMode(root *TreeNode) []int {
	maxCount := 0
	maxNum := []int{}
	count := 0
	var pre *TreeNode
	var mid func(cur *TreeNode)
	mid = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		mid(cur.Left)
		if pre == nil {
			count = 1
		} else if pre.Val == cur.Val {
			count++
		} else {
			count = 1
		}
		pre = cur
		if count == maxCount {
			maxNum = append(maxNum, cur.Val)
		} else if count > maxCount {
			maxCount = count
			maxNum = []int{cur.Val}
		}
		mid(cur.Right)
	}
	mid(root)
	return maxNum
}
