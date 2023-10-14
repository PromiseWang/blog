package main

import (
	"fmt"
	"leetcode/6_tree/day21"
)

func pre(root *day21.TreeNode) {
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
	var root1 *day21.TreeNode
	root1 = day21.Init4(root1)
	result := day21.T(root1, 3)
	pre(result)
}
