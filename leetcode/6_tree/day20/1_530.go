package day20

import "math"

// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 差值是一个正数，其数值等于两值之差的绝对值。
func getMinimumDifference(root *TreeNode) int {
	var path []int
	inorder(root, &path)
	minNum := math.MaxInt32
	for i := 1; i < len(path); i++ {
		if path[i]-path[i-1] < minNum {
			minNum = path[i] - path[i-1]
		}
	}
	return minNum
}

func inorder(root *TreeNode, path *[]int) {
	if root.Left != nil {
		inorder(root.Left, path)
	}
	*path = append(*path, root.Val)
	if root.Right != nil {
		inorder(root.Right, path)
	}
}

func getMinimumDifference2(root *TreeNode) int {
	minNum := math.MaxInt32
	var pre *TreeNode
	var inorder2 func(cur *TreeNode)
	inorder2 = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		inorder2(cur.Left)
		if pre != nil && cur.Val-pre.Val < minNum {
			minNum = cur.Val - pre.Val
		}
		pre = cur
		inorder2(cur.Right)
	}
	inorder2(root)
	return minNum
}
