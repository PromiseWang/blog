package main

import (
	"fmt"
	"leetcode/6_tree/day17"
)

func main() {
	//root := &day17.TreeNode{
	//	Val:   1,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Left = &day17.TreeNode{
	//	Val:   2,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Right = &day17.TreeNode{
	//	Val:   3,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Left.Left = &day17.TreeNode{
	//	Val:   4,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Left.Right = &day17.TreeNode{
	//	Val:   5,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Right.Left = &day17.TreeNode{
	//	Val:   6,
	//	Left:  nil,
	//	Right: nil,
	//}
	root := &day17.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left = &day17.TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	root.Right = &day17.TreeNode{
		Val:   20,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = &day17.TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	root.Right.Right = &day17.TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	result := day17.T(root)
	fmt.Println(result)

}
