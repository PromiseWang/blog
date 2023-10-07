package day14

// 每一层最大值

import (
	"github.com/emirpasic/gods/queues/arrayqueue"
	"math"
)

func largestValues(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		maxNum := math.MinInt32
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			if node.(*TreeNode).Val > maxNum {
				maxNum = node.(*TreeNode).Val
			}
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
		}
		result = append(result, maxNum)
	}
	return result
}
