package day07

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	map4 := map[int]int{}
	for _, v := range nums4 {
		map4[v]++
	}
	for _, i := range nums1 {
		for _, j := range nums2 {
			for _, k := range nums3 {
				if v, ok := map4[0-i-j-k]; ok {
					count += v
				}
			}
		}
	}
	return count
}


func fourSumCount1(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	map1 := map[int]int{}
	map2 := map[int]int{}
	map3 := map[int]int{}
	map4 := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]]++
		map2[nums2[i]]++
		map3[nums3[i]]++
		map4[nums4[i]]++
	}
	for k1, v1 := range map1 {
		for k2, v2 := range map2 {
			for k3, v3 := range map3 {
				if v4, ok := map4[0-k1-k2-k3]; ok {
					count += v1 * v2 * v3 * v4
				}
			}
		}
	}
	return count
}

func fourSumCount2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	mapAB := map[int]int{}
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			mapAB[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			if v, ok := mapAB[0-v3-v4]; ok {
				count += v
			}
		}
	}
	return count
}