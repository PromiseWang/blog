package day07

import (
	"reflect"
	"sort"
)

func backTracking(nums []int, index int, result *[][]int, path *[]int) {
	if len(*path) == 3 {
		if (*path)[0]+(*path)[1]+(*path)[2] == 0 {
			// 后面的append()操作是深拷贝，所以这里使用temp保存path的值
            temp := make([]int, len(*path))  
			copy(temp, *path)
			for _, v := range *result {
				if reflect.DeepEqual(temp, v) {  // 使用反射判断temp是否在result中
					return
				}
			}
			*result = append(*result, temp)
		}
		return
	}
	for i := index; i < len(nums); i++ {
		*path = append(*path, nums[i])
		backTracking(nums, i+1, result, path)
		*path = (*path)[:len(*path)-1]

	}
}

func threeSum(nums []int) [][]int {
	var result [][]int
	var path []int
	sort.Ints(nums)
	backTracking(nums, 0, &result, &path)
	return result
}


func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	map1 := map[int]int{}
	for _, v := range nums {
		map1[v]++
	}
	var result [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if v, ok := map1[-nums[i]-nums[j]]; ok {
				temp := []int{nums[i], nums[j], -nums[i] - nums[j]}
				sort.Ints(temp)
				count := 0
				if nums[i] == nums[j] || nums[j] == -nums[i]-nums[j] || nums[i] == -nums[i]-nums[j] {
					count = 2
				}
				if nums[i] == nums[j] && nums[j] == -nums[i]-nums[j] {
					count = 3
				}
				if v >= count {
					flag := true
					for _, value := range result {
						if reflect.DeepEqual(value, temp) {
							flag = false
						}
					}
					if flag {
						result = append(result, temp)
					}

				}
			}
		}
	}
	return result
}


func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}