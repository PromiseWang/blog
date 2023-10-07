package day14

import "github.com/emirpasic/gods/queues/arrayqueue"

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxNum(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 层序
func maxDepth1(root *TreeNode) int {
	var result [][]int
	if root == nil {
		return 0
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		temp := []int{}
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			temp = append(temp, node.(*TreeNode).Val)
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
		}
		result = append(result, temp)
	}
	return len(result)
}

func MaxDepth(root *TreeNode) int {
	return maxDepth(root)
}
