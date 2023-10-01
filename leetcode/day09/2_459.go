package day09

import "strings"

func repeatedSubstringPattern(s string) bool {
	newString := s + s
	newString = newString[1 : len(newString)-2]
	return strings.Contains(newString, s)
}