package day08

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func reverseStr(s string, k int) string {
	newString := ""
	for i := 0; i < len(s); i += 2 * k {
		for j := min(len(s), i+k) - 1; j >= i; j-- { // 前k段
			newString += string(s[j])
		}
		for j := i + k; j < min(len(s), i+2*k); j++ {  // 后k段
			newString += string(s[j])
		}
	}
	return newString
}