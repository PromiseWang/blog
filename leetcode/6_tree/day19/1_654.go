package day19

// 给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:
//
// 创建一个根节点，其值为 nums 中的最大值。
// 递归地在最大值 左边 的 子数组前缀上 构建左子树。
// 递归地在最大值 右边 的 子数组后缀上 构建右子树。
// 返回 nums 构建的 最大二叉树 。
func constructMaximumBinaryTree(nums []int) *TreeNode {
	max := 0
	maxIndex := 0
	for i, v := range nums {
		if v > max {
			max = v
			maxIndex = i
		}
	}
	root := &TreeNode{
		Val:   max,
		Left:  nil,
		Right: nil,
	}
	if len(nums) == 0 {
		return root
	}
	leftNums := nums[:maxIndex]
	rightNums := nums[maxIndex+1:]
	if len(leftNums) != 0 {
		root.Left = constructMaximumBinaryTree(leftNums)
	}
	if len(rightNums) != 0 {
		root.Right = constructMaximumBinaryTree(rightNums)
	}
	return root
}
