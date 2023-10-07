package day15

import "github.com/emirpasic/gods/queues/arrayqueue"

// n叉树最大深度
func nMaxDepth(root *Node) int {
	var result [][]int
	if root == nil {
		return 0
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		var temp []int
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			temp = append(temp, node.(*Node).Val)
			for j := 0; j < len(node.(*Node).Children); j++ {
				if node.(*Node).Children[j] != nil {
					queue.Enqueue(node.(*Node).Children[j])
				}
			}
		}
		result = append(result, temp)
	}
	return len(result)
}
func maxNum1(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func recursionMaxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 1
	for _, v := range root.Children {
		depth = maxNum1(depth, recursionMaxDepth(v)+1)
	}
	return depth
}
