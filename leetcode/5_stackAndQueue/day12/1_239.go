package day12

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