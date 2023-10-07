package day10

import (
	"github.com/emirpasic/gods/stacks/arraystack"
)

type MyQueue struct {
	Stack1 arraystack.Stack
	Stack2 arraystack.Stack
}

func Constructor() MyQueue {
	return MyQueue{
		Stack1: *arraystack.New(),
		Stack2: *arraystack.New(),
	}
}

func (this *MyQueue) Push(x int) {
	this.Stack1.Push(x)
}

func (this *MyQueue) Pop() int {
	for !this.Stack1.Empty() {
		if v, ok := this.Stack1.Pop(); ok {
			this.Stack2.Push(v)
		}
	}
	value, _ := this.Stack2.Pop()
	for !this.Stack2.Empty() {
		if v, ok := this.Stack2.Pop(); ok {
			this.Stack1.Push(v)
		}
	}
	return value.(int)
}

func (this *MyQueue) Peek() int {
	for !this.Stack1.Empty() {
		if v, ok := this.Stack1.Pop(); ok {
			this.Stack2.Push(v)
		}
	}
	value, _ := this.Stack2.Peek()
	for !this.Stack2.Empty() {
		if v, ok := this.Stack2.Pop(); ok {
			this.Stack1.Push(v)
		}
	}
	return value.(int)
}

func (this *MyQueue) Empty() bool {
	return this.Stack1.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// ["MyQueue", "push", "push", "peek", "pop", "empty"]
//[[], [1], [2], [], [], []]
