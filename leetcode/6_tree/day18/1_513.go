package day18

import (
	"fmt"
	"github.com/emirpasic/gods/queues/arrayqueue"
)

// 找这棵树最左下角的值
func findBottomLeftValue(root *TreeNode) int {
	queue := arrayqueue.New()
	var result [][]int
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
	fmt.Println(result)
	return result[len(result)-1][0]
}
