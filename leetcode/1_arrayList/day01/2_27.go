package day01

// 方法一
func removeElement(nums []int, val int) int {
	n := len(nums)
	left := 0
	for right := 0; right < n; right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

// 方法二
func removeElement2(nums []int, val int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		for left <= right {
			if nums[left] == val {
				break
			} else {
				left++
			}
		}
		for left <= right {
			if nums[right] != val {
				break
			} else {
				right--
			}
		}
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	return left
}