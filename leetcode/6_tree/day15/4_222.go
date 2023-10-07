package day15

import "github.com/emirpasic/gods/queues/arrayqueue"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	count := 0
	for !queue.Empty() {
		size := queue.Size()
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			count++
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
		}
	}
	return count
}

func T(root *TreeNode) int {
	return countNodes(root)
}
