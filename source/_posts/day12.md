---
title: 代码随想录算法训练营第十二天 239.滑动窗口最大值、347.前K个高频元素。
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 栈与队列
categories: 刷题
abbrlink: 9ab3b9d2
date: 2023-10-04 19:45:02
---

# 代码随想录算法训练营第十二天 239.滑动窗口最大值、347.前K个高频元素。

## 239.滑动窗口最大值

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/sliding-window-maximum/)
>
>   文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0239.%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%9C%80%E5%A4%A7%E5%80%BC.html)
>
>   视频讲解：[单调队列正式登场！| LeetCode：239. 滑动窗口最大值](https://www.bilibili.com/video/BV1XS4y1p7qj)
>
>   状态：TLE

### 思路

#### 方法一

遍历每个滑动窗口（看作一个队列），升序排序后取最大（复制一段新的数组，在新的里面排序）。时间复杂度为O(n^2log n)。超时

#### 方法二（卡哥讲解）

1.   准备一个自定义的双向队列数据结构，让其中的`front`端保持最大值，并且重构`Pop()`和`Push()`操作。
2.   先将前K的元素`Push()`入队，但是要进行判断。如果后入队的元素大于之前入队的元素，那么之前的元素`PopBack()`。可以保持`front`端一直为最大。
3.   `Pop()`除了最大值的所有元素
4.   `Push()`一个新的元素
5.   `GetMaxValue()`，并且加入到结果切片中
6.   重复3\~5直至`nums`结尾

### 代码

#### 重构两个函数，新增功能GetMaxValue()

-   `GetMaxValue()`：会保持在队列`front`的位置，可以使用`Front()`操作取代。

``` go
func (this *Deque) Front() int {
	return this.queue[0]
}
```

-   `Pop()`：出队时，只有一种情况，旧的最大值不在新的滑动窗口内，旧的最大值出队（切片舍去第一个元素）

``` go
func (this *Deque) Pop(value int) {
	if !this.Empty() && value == this.Front() {
		this.queue = this.queue[1:]
	}
}
```

-   `Push()`：入队时，如果入队的值`value`大于队列`back`位置的元素，让该元素`PopBack()`（切片舍去最后一个元素）

``` go
func (this *Deque) Push(value int) {
	for !this.Empty() && value > this.Back() {
		this.queue = this.queue[:len(this.queue)-1]
	}
	this.queue = append(this.queue, value)
}
```

### 代码

``` go
type Deque struct {
	queue []int
}

func Constructor() *Deque {
	return &Deque{queue: make([]int, 0)}
}

func (this *Deque) Front() int {
	return this.queue[0]
}

func (this *Deque) Back() int {
	return this.queue[len(this.queue)-1]
}

func (this *Deque) Empty() bool {
	return len(this.queue) == 0
}

func (this *Deque) Push(value int) {
	for !this.Empty() && value > this.Back() {
		this.queue = this.queue[:len(this.queue)-1]
	}
	this.queue = append(this.queue, value)
}

func (this *Deque) Pop(value int) {
	if !this.Empty() && value == this.Front() {
		this.queue = this.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	deque := Constructor()
	var result []int
	for i := 0; i < k; i++ {
		deque.Push(nums[i])
	}
	result = append(result, deque.Front())
	for i := k; i < len(nums); i++ {
		deque.Pop(nums[i-k])
		deque.Push(nums[i])
		result = append(result, deque.Front())
	}
	return result
}
```

## 347. 前K个高频元素

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/top-k-frequent-elements/)
>
>   文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0347.%E5%89%8DK%E4%B8%AA%E9%AB%98%E9%A2%91%E5%85%83%E7%B4%A0.html)
>
>   视频讲解：[优先级队列正式登场！大顶堆、小顶堆该怎么用？| LeetCode：347.前 K 个高频元素](https://www.bilibili.com/video/BV1Xg41167Lz)
>
>   状态：AC

### 思路

1.   构建键值对<key, value>，`key`为数字，`value`为该数字出现的次数。
2.   生成优先队列，优先级为`value`
3.   取优先队列的前K个元素

### 代码

``` go
// 优先队列中的条目
type Element struct {
	Num     int
	Count   int
}
// 比较器函数，根据Count降序
func byPriority(a, b interface{}) int {
	priorityA := a.(Element).Count
	priorityB := b.(Element).Count
	return -utils.IntComparator(priorityA, priorityB) // "-" 代表降序
}

func topKFrequent(nums []int, k int) []int {
	// 构建键值对
    maps := map[int]int{}
	for _, v := range nums {
		maps[v]++
	}
    // 初始化优先队列
	pq := priorityqueue.NewWith(byPriority)
	for k, v := range maps {
		pq.Enqueue(Element{
			Num:     k,
			Count: v,
		})
	}
	var result []int
    // 取前k个元素保存到result中
	for i := 0; i < k; i++ {
		value, _ := pq.Dequeue()
		result = append(result, value.(Element).Num)
	}
	return result
}
```

