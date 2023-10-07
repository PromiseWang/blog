package day13

import "github.com/emirpasic/gods/stacks/arraystack"

func PostorderTraversal(root *TreeNode) []int {
	var result []int
	getResultPostorder(root, &result)
	return result
}

func getResultPostorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	getResultPostorder(root.Left, result)
	getResultPostorder(root.Right, result)
	*result = append(*result, root.Val)
}

func PostorderTraversal1(root *TreeNode) []int {
	stack := arraystack.New()
	result := []int{}
	if root == nil {
		return result
	}
	stack.Push(root)
	for !stack.Empty() {
		value, _ := stack.Pop()
		result = append(result, value.(*TreeNode).Val)
		if value.(*TreeNode).Left != nil {
			stack.Push(value.(*TreeNode).Left)
		}
		if value.(*TreeNode).Right != nil {
			stack.Push(value.(*TreeNode).Right)
		}
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}
