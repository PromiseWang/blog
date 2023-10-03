package day11

import "github.com/emirpasic/gods/stacks/arraystack"

func IsValid(s string) bool {
	stack := arraystack.New()
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack.Push(v)
		case ')':
			temp, _ := stack.Pop()
			if temp == '(' {
				continue
			} else {
				return false
			}
		case ']':
			temp, _ := stack.Pop()
			if temp == '[' {
				continue
			} else {
				return false
			}
		case '}':
			temp, _ := stack.Pop()
			if temp == '{' {
				continue
			} else {
				return false
			}
		}
	}
	return stack.Empty()
}
