package day17

import (
	"strconv"
)

// 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。
func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	path := []int{}
	getResult(root, &result, &path)
	return result
}

func getResult(root *TreeNode, result *[]string, path *[]int) {
	if root.Left == nil && root.Right == nil {
		*path = append(*path, root.Val)
		var temp string
		if len(*path) == 1 {
			temp = "1"
		} else {
			temp = strconv.Itoa((*path)[0])
			for i := 1; i < len(*path); i++ {
				temp += "->" + strconv.Itoa((*path)[i])
			}
		}
		*path = (*path)[:len(*path)-1]
		*result = append(*result, temp)
		return
	}
	if root.Left != nil {
		*path = append(*path, root.Val)
		getResult(root.Left, result, path)
		*path = (*path)[:len(*path)-1]
	}
	if root.Right != nil {
		*path = append(*path, root.Val)
		getResult(root.Right, result, path)
		*path = (*path)[:len(*path)-1]
	}
}
