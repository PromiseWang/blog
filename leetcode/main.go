package main

import (
	"fmt"
	"leetcode/6_tree/day20"
)

func pre(root *day20.TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	if root.Left != nil {
		pre(root.Left)
	}
	if root.Right != nil {
		pre(root.Right)
	}
}

func main() {
	var root1 *day20.TreeNode
	root1 = day20.Init4(root1)
	p := &day20.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	q := &day20.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	result := day20.T(root1, p, q)
	fmt.Println(result.Val)

}
