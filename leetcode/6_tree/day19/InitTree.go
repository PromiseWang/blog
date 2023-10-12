package day19

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	    1
//	  /  \
//	 2    3
//	/ \  / \
//
// 4  5 6  7
func Init(root *TreeNode) *TreeNode {
	root = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left.Left = &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right = &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	root.Right.Right = &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	return root
}

//	    4
//	  /   \
//	 2     7
//	/ \
//
// 1   3
func Init2(root *TreeNode) *TreeNode {
	root = &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	root.Left.Left = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	return root
}

//		    62
//		  /   \
//		 2     93
//		  \
//	     3
//	    /
//	   15
func Init3(root *TreeNode) *TreeNode {
	root = &TreeNode{
		Val:   62,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   93,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right = &TreeNode{
		Val:   30,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right.Left = &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	return root
}

func Init4(root *TreeNode) *TreeNode {
	root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	return root
}
