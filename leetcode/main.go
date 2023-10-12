package main

import (
	"fmt"
	"leetcode/6_tree/day19"
)

func pre(root *day19.TreeNode) {
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
	var root1 *day19.TreeNode
	root1 = day19.Init4(root1)
	result := day19.T(root1)
	fmt.Println(result)

}
