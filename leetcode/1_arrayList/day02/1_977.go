package day02

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left := 0
	right := len(nums) - 1
	for i := right; i >= 0; i-- {  // i为result的索引
		if nums[left]*nums[left] < nums[right]*nums[right] {  // right侧更大
			result[i] = nums[right] * nums[right]
			right--
		} else {  // left侧更大
			result[i] = nums[left] * nums[left]
			left++
		}
	}
	return result
}
