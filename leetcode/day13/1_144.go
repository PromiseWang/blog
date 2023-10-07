package day13

func preorderTraversal(root *TreeNode) []int {
	var result []int
	p := root
	getResult(p, &result)
	return result
}

func getResult(root *TreeNode, result *[]int) {
	for root != nil {
		*result = append(*result, root.Val)
		getResult(root.Left, result)
		getResult(root.Right, result)
	}
}
