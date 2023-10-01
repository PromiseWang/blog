package main

import (
	"fmt"
	"leetcode/day10"
)

// ["MyStack", "push", "push", "top", "pop", "empty"]
// [[], [1], [2], [], [], []]
func main() {
	obj := day10.Constructor1() // MyQueue
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
