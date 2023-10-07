package day02

func generateMatrix(n int) [][]int {
	top := 0
	bottom := n - 1
	left := 0
	right := n - 1
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	num := 1
	for i := 0; i < n/2; i++ {
		for j := left; j < right; j++ {
			result[top][j] = num
			num++
		}

		for j := top; j < bottom; j++ {
			result[j][right] = num
			num++
		}

		for j := right; j > left; j-- {
			result[bottom][j] = num
			num++
		}
		for j := bottom; j > top; j-- {
			result[j][left] = num
			num++
		}
		top++
		left++
		right--
		bottom--
	}
	if n%2 == 1 {
		result[n/2][n/2] = num
	}
	return result
}