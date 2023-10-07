package day14

import (
	"github.com/emirpasic/gods/queues/arrayqueue"
	"math"
)

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minN := math.MaxInt32
	if root.Left != nil {
		minN = minNum(minDepth(root.Left), minN)
	}
	if root.Right != nil {
		minN = minNum(minDepth(root.Right), minN)
	}
	return minN + 1
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	count := 0
	for !queue.Empty() {
		count++
		size := queue.Size()
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
			if node.(*TreeNode).Left == nil && node.(*TreeNode).Right == nil {
				return count
			}
		}
	}
	return count
}

func MinDepth(root *TreeNode) int {
	return minDepth(root)
}
