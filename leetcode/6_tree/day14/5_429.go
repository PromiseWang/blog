package day14

import "github.com/emirpasic/gods/queues/arrayqueue"

type Node1 struct {
	Val      int
	Children []*Node1
}

func levelOrder(root *Node1) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		temp := []int{}
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			temp = append(temp, node.(*Node1).Val)
			for j := 0; j < len(node.(*Node1).Children); j++ {
				if node.(*Node1).Children[j] != nil {
					queue.Enqueue(node.(*Node1).Children[j])
				}
			}
		}
		result = append(result, temp)
	}
	return result
}
