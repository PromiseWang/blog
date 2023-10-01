package day05
//遍历两次
func twoSum(nums []int, target int) []int {
    mapNums := map[int]int{}
    for i, v := range nums {
       mapNums[v] = i
    }
    var result []int
    for i, v := range nums {
       if _, ok := mapNums[target-v]; ok && mapNums[target-v] != i {
          result = append(result, i)
          result = append(result, mapNums[target-v])
          break
       }
    }
    return result
}

func twoSum1(nums []int, target int) []int {
	mapNums := map[int]int{}
	for i, v := range nums {
		if value, ok := mapNums[target-v]; ok {
			return []int{i, value}
		} else {
			mapNums[v] = i
		}
	}
	return []int{}
}