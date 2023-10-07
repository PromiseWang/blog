package day05

const (
	NEW = iota
	EXIST1
	EXIST2
)

func intersection(nums1 []int, nums2 []int) []int {
	set := map[int]int{}  //类集合操作
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = NEW  // 第一个出现
		} else {  // 多次出现
			set[v] = EXIST1
		}
	}
	for _, v := range nums2 {
		if _, ok := set[v]; ok {  // nums2也出现
			set[v] = EXIST2
		}
	}
	var result []int
	for k, v := range set {
		if v == EXIST2 {
			result = append(result, k)
		}
	}
	return result
}