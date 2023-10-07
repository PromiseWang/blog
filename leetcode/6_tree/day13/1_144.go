package day13

import (
	"github.com/emirpasic/gods/stacks/arraystack"
)

// 递归前序遍历
func preorderTraversal(root *TreeNode) []int {
	var result []int
	getResultPreorder(root, &result)
	return result
}

func getResultPreorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	getResultPreorder(root.Left, result)
	getResultPreorder(root.Right, result)
}

// 非递归前序遍历
func preorderTraversal1(root *TreeNode) []int {
	stack := arraystack.New()
	result := []int{}
	if root == nil {
		return result
	}
	stack.Push(root)
	for !stack.Empty() {
		value, _ := stack.Pop()
		result = append(result, value.(*TreeNode).Val)
		if value.(*TreeNode).Right != nil {
			stack.Push(value.(*TreeNode).Right)
		}
		if value.(*TreeNode).Left != nil {
			stack.Push(value.(*TreeNode).Left)
		}
	}
	return result
}
