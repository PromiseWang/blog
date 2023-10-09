package main

import (
	"fmt"
	"leetcode/6_tree/day15"
)

func main() {
	root := &day15.TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left = &day15.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &day15.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left.Left = &day15.TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right = &day15.TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = &day15.TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	result := day15.T(root)
	fmt.Println(result)
}
