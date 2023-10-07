package main

import (
	"leetcode/day13"
)

func main() {
	root := &day13.TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	day13.InsertRight(2, root)
	day13.InsertLeft(3, root)
}
