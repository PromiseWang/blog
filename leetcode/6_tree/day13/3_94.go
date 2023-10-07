package day13

import "github.com/emirpasic/gods/stacks/arraystack"

func inorderTraversal(root *TreeNode) []int {
	var result []int
	getResultInorder(root, &result)
	return result
}

func getResultInorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	getResultInorder(root.Left, result)
	*result = append(*result, root.Val)
	getResultInorder(root.Right, result)
}

func InorderTraversal1(root *TreeNode) []int {
	stack := arraystack.New()
	result := []int{}
	if root == nil {
		return result
	}
	p := root
	for !stack.Empty() || p != nil {
		if p != nil {
			stack.Push(p)
			p = p.Left
		} else {
			node, _ := stack.Pop()
			p = node.(*TreeNode)
			result = append(result, p.Val)
			p = p.Right
		}
	}
	return result
}
