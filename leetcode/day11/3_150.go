package day11

import (
	"github.com/emirpasic/gods/stacks/arraystack"
	"strconv"
)

func EvalRPN(tokens []string) int {
	stack := arraystack.New()
	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err == nil {
			stack.Push(num)
		} else {
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			if v == "+" {
				stack.Push(a.(int) + b.(int))
			} else if v == "-" {
				stack.Push(a.(int) - b.(int))
			} else if v == "*" {
				stack.Push(a.(int) * b.(int))
			} else if v == "/" {
				stack.Push(a.(int) / b.(int))
			}
		}
	}
	result, _ := stack.Peek()
	return result.(int)
}
