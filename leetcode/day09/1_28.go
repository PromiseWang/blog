package day09

func getNext(next []int, s string) {
	j := 0
	next[0] = 0
	for i := 1; i < len(s); i++ {
		for s[i] != s[j] && j > 0 {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
		//fmt.Printf("第%v次循环后的next数组结果为:%v\n", i, next)
	}
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle))
	getNext(next, needle)
	j := 0
	for i := range haystack {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}