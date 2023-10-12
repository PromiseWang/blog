package day18

// 路径总和
//给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点的路径，
//这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
//
//叶子节点 是指没有子节点的节点。

func hasPathSum(root *TreeNode, targetSum int) bool {
	var result [][]int
	var path []int
	getResult(root, &result, &path)
	for i := 0; i < len(result); i++ {
		sum := 0
		for _, v := range result[i] {
			sum += v
		}
		if sum == targetSum {
			return true
		}
	}
	return false
}

func getResult(root *TreeNode, result *[][]int, path *[]int) {
	if root.Left == nil && root.Right == nil {
		*path = append(*path, root.Val)
		temp := make([]int, len(*path))
		copy(temp, *path)
		*result = append(*result, temp)
		*path = (*path)[:len(*path)-1]
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
