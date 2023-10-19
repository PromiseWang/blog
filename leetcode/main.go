package main

import (
	"fmt"
	"leetcode/7_backtracking/day26"
)

func main() {
	//a := []int{2, 3, 6, 7}
	a := []int{2, 3, 5}
	b := 8
	result := day26.T(a, b)
	fmt.Println(result)
}
