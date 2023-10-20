package day26

import (
	"sort"
	"strconv"
)

// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用 一次 。
//
// 注意：解集不能包含重复的组合。
func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var path []int
	sort.Ints(candidates)
	maps := map[string]bool{}
	var backTracking func(candidates []int, target, sum, index int)
	backTracking = func(candidates []int, target, sum, index int) {
		if target == sum {
			s := ""
			for _, v := range path {
				s += strconv.Itoa(v)
			}
			if !maps[s] {
				temp := make([]int, len(path))
				copy(temp, path)
				result = append(result, temp)
				maps[s] = true
			}
			return
		}
		for i := index; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && len(path) == 0 {
				continue
			}
			if sum+candidates[i] > target {
				return
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backTracking(candidates, target, sum, i+1)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backTracking(candidates, target, 0, 0)
	return result
}

func T(a []int, b int) [][]int {
	return combinationSum2(a, b)
}
