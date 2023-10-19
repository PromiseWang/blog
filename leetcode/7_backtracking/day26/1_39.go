package day26

import (
	"sort"
	"strconv"
)

// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
//
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
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
	return combinationSum(a, b)
}
