package day05

import "reflect"

//麻烦的方法
func isAnagram(s string, t string) bool {
	words1 := map[byte]int{}
	words2 := map[byte]int{}
	for i := 0; i < len(s); i++ {
		_, ok := words1[s[i]]
		if ok {
			words1[s[i]] += 1
		} else {
			words1[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		_, ok := words2[t[i]]
		if ok {
			words2[t[i]] += 1
		} else {
			words2[t[i]] = 1
		}
	}

	return reflect.DeepEqual(words1, words2)
}

// 简单的代码
func isAnagram1(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}