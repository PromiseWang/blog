package day01

func searchRange(nums []int, target int) []int {
	result := []int{}
    var temp int
    temp = binary(nums, target, true)
    if temp == -1 {  // 如果查找开始元素为-1，说明该数不存在，第二次二分不需要进行
        return []int{-1, -1}
    } else {
		result = append(result, temp)
        result = append(result, binary(nums, target, false))
        return result
    }
}

func binary(nums []int, target int, flag bool) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if flag {
			if nums[mid] == target && (mid == 0 || nums[mid-1] < target) {
				return mid
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] == target && (mid == len(nums)-1 || nums[mid+1] > target) {
				return mid
			} else if nums[mid] <= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}