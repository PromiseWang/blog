package day02

func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	sum := 0
	minLength := 100001
	length := len(nums)
	for left <= right && right <= length {
		if sum < target {
			if right < length {
				sum += nums[right]
				right++
			} else {
				break
			}
		} else {
			if minLength > right-left {
				minLength = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	if minLength == 100001 {
		return 0
	}
	return minLength
}