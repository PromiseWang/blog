package main

import (
	"fmt"
	"leetcode/day11"
)

func main() {
	s := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	fmt.Println(day11.EvalRPN(s))
}
