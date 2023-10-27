---
title: 代码随想录算法训练营第十天 232.用栈实现队列、225. 用队列实现栈。
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 栈与队列
categories: 刷题
abbrlink: a8d0ca77
date: 2023-10-02 00:36:27
---

# 代码随想录算法训练营第十天 232.用栈实现队列、225. 用队列实现栈。

## 232.用栈实现队列

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/implement-queue-using-stacks/)
>
> 文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0232.%E7%94%A8%E6%A0%88%E5%AE%9E%E7%8E%B0%E9%98%9F%E5%88%97.html)
>
> 视频讲解：[栈的基本操作！ | LeetCode：232.用栈实现队列](https://www.bilibili.com/video/BV1nY4y1w7VC)
>
> 状态：AC

### 思路

- 初始化两个栈，一个作为主栈`Stack1`、一个作为副栈`Stack2`
- 主栈存放数据、副栈临时放数据
- `Push`操作：将数据放入`Stack1`中
- `Pop`操作：
    - 将`Stack1`中所有元素退栈，并入栈`Stack2`。
    - 将`Stack2`的栈顶元素出栈
    - 将`Stack2`中所有元素退栈，并入栈`Stack1`。
- `Peek`操作：
    - 将`Stack1`中所有元素退栈，并入栈`Stack2`。
    - 取`Stack2`的栈顶元素
    - 将`Stack2`中所有元素退栈，并入栈`Stack1`。
- `Empty`操作：
    - 查看`Stack1`是否为空即可

### 代码

``` go
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
```

## 225. 用队列实现栈

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/implement-stack-using-queues/)
>
> 文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0225.%E7%94%A8%E9%98%9F%E5%88%97%E5%AE%9E%E7%8E%B0%E6%A0%88.html)
>
> 视频讲解：[队列的基本操作！ | LeetCode：225. 用队列实现栈](https://www.bilibili.com/video/BV1Fd4y1K7sm)
>
> 状态：AC

### 思路

- 初始化一个队列`Queue1`
- 主栈存放数据、副栈临时放数据
- `Push`操作：将数据入队`Queue1`
- `Pop`操作：
    - 将`Queue1`前`n-1`个元素出队，随后入队在队尾
    - 最后一个元素出队，返回
- `Top`操作：
    - 将`Queue`每个出队，随后入队在队尾
    - 如果是最后一个元素出队，保存该值，并入队
    - 返回那个值
- `Empty`操作：
    - 查看`Queue1`是否为空即可

### 代码

``` go
type MyStack struct {
	Queue1 arrayqueue.Queue
}

func Constructor() MyStack {
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

```

