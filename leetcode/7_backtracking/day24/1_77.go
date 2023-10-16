package day24

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。
func combine(n int, k int) [][]int {
	path := []int{}
	result := [][]int{}
	var backTracking func(n, k, index int)
	backTracking = func(n, k, index int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := index; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backTracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	backTracking(n, k, 1)
	return result
}

func T(n, k int) [][]int {
	return combine(n, k)
}
