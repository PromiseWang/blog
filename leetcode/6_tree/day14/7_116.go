package day14

import "github.com/emirpasic/gods/queues/arrayqueue"

// 与117题代码相同
// 二叉树向右链接节点

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := arrayqueue.New()
	root.Next = nil
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		var p *Node = nil
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()

			if node.(*Node).Left != nil {
				queue.Enqueue(node.(*Node).Left)
			}
			if node.(*Node).Right != nil {
				queue.Enqueue(node.(*Node).Right)
			}
			if p == nil {
				p = node.(*Node)
			} else {
				p.Next = node.(*Node)
				p = p.Next
			}
		}
		p.Next = nil
	}
	return root
}
