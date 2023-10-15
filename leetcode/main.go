package main

import (
	"fmt"
	"leetcode/6_tree/day22"
)

func pre(root *day22.TreeNode) {
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
	var root1 *day22.TreeNode
	root1 = day22.Init(root1)
	result := day22.T(root1)
	pre(result)
}
