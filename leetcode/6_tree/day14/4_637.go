package day14

import "github.com/emirpasic/gods/queues/arrayqueue"

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	if root == nil {
		return result
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		sum := 0
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			sum += node.(*TreeNode).Val
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
		}
		result = append(result, float64(sum)/float64(size))
	}
	return result
}

func AverageOfLevels(root *TreeNode) []float64 {
	return averageOfLevels(root)
}
