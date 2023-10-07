package day10

import "github.com/emirpasic/gods/queues/arrayqueue"

type MyStack struct {
	Queue1 arrayqueue.Queue
}

func Constructor1() MyStack {
	return MyStack{
		Queue1: *arrayqueue.New(),
	}
}

func (this *MyStack) Push(x int) {
	this.Queue1.Enqueue(x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < this.Queue1.Size()-1; i++ {
		v, _ := this.Queue1.Dequeue()
		value := v.(int)
		this.Queue1.Enqueue(value)
	}
	v, _ := this.Queue1.Dequeue()
	value := v.(int)
	return value
}

func (this *MyStack) Top() int {
	top := 0
	for i := 0; i < this.Queue1.Size(); i++ {
		v, _ := this.Queue1.Dequeue()
		value := v.(int)
		this.Queue1.Enqueue(value)
		if i == this.Queue1.Size()-1 {
			top = value
		}
	}
	return top
}

func (this *MyStack) Empty() bool {
	return this.Queue1.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
