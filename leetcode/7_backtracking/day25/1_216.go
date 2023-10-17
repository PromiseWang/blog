package day25

// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
//
// 只使用数字1到9
// 每个数字 最多使用一次
// 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
func combinationSum3(k int, n int) [][]int {
	var result [][]int
	var path []int
	var backTracking func(k, n, index, sum int)
	backTracking = func(k, n, index, sum int) {
		if len(path) == k {
			if sum == n {
				temp := make([]int, k)
				copy(temp, path)
				result = append(result, temp)
			}
			return
		}
		for i := index; i <= 9; i++ {
			if sum+i > n {
				return
			}
			path = append(path, i)
			sum += i
			backTracking(k, n, i+1, sum)
			sum -= path[len(path)-1]
			path = path[:len(path)-1]
		}
	}
	backTracking(k, n, 1, 0)
	return result
}
