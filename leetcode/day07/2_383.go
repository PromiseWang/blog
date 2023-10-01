package day07

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	arr := make([]int, 26)
	for _, v := range magazine {
		arr[v-'a']++
	}
	for _, v := range ransomNote {
		arr[v-'a']--
		if arr[v-'a'] < 0 {
			return false
		}
	}
	return true
}