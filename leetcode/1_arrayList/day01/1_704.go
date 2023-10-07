package day01


func search(nums []int, target int) int {
	// 左闭右闭写法
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func search1(nums []int, target int) int {
	// 左闭右开写法
	left := 0
	right := len(nums)
	var mid int
	for left < right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return -1
}