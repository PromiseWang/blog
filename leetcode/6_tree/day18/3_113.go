package day18

//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
//叶子节点 是指没有子节点的节点。

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var path []int
	getResult1(root, &result, &path, targetSum)
	return result
}

func getResult1(root *TreeNode, result *[][]int, path *[]int, target int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*path = append(*path, root.Val)
		sum := 0
		for _, v := range *path {
			sum += v
		}
		if sum == target {
			temp := make([]int, len(*path))
			copy(temp, *path)
			*result = append(*result, temp)
		}
		*path = (*path)[:len(*path)-1]
		return
	}
	if root.Left != nil {
		*path = append(*path, root.Val)
		getResult1(root.Left, result, path, target)
		*path = (*path)[:len(*path)-1]
	}
	if root.Right != nil {
		*path = append(*path, root.Val)
		getResult1(root.Right, result, path, target)
		*path = (*path)[:len(*path)-1]
	}
}

func T(root *TreeNode, targetSum int) [][]int {
	return pathSum(root, targetSum)
}
