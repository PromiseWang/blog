package day22

// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
//
// 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := backTracking(nums, 0, len(nums)-1)
	return root
}

func backTracking(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  backTracking(nums, left, mid-1),
		Right: backTracking(nums, mid+1, right),
	}
	return root
}
