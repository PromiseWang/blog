package day25

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
func letterCombinations(digits string) []string {
	var result []string
	var path string
	words := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}
	var backTracking func(index int)
	backTracking = func(index int) {
		if len(path) == len(digits) && len(path) != 0 {
			result = append(result, path)
			return
		}
		for i1 := index; i1 < len(digits); i1++ {
			for i2 := 0; i2 < len(words[int(digits[i1]-50)]); i2++ {
				path += words[int(digits[i1]-50)][i2]
				backTracking(i1 + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backTracking(0)

	return result
}

func T(s string) []string {
	return letterCombinations(s)
}
