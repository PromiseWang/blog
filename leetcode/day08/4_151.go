package day08

func reverseWords(s string) string {
	words := []string{}
	left := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if s[left:i] != "" {
				words = append(words, s[left:i])
			}
			left = i + 1
		} else if i == len(s)-1 {
			words = append(words, s[left:])
		}
	}
	newString := ""
	for i := len(words) - 1; i >= 0; i-- {
		newString += words[i]
		if i > 0 {
			newString += " "
		}
	}
	return newString
}