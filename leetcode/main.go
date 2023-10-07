package main

import (
	"fmt"
	"leetcode/6_tree/day14"
)

func main() {
	root := &day14.TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left = &day14.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &day14.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Left.Left = &day14.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	root.Right.Left = &day14.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	result := day14.T(root)
	fmt.Println(result)
}
