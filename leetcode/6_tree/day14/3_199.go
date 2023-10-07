package day14

import "github.com/emirpasic/gods/queues/arrayqueue"

func RightSideView(root *TreeNode) []int {
	var results [][]int
	if root == nil {
		return []int{}
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
		results = append(results, temp)
	}
	result := []int{}
	for _, v := range results {
		result = append(result, v[len(v)-1])
	}
	return result
}
