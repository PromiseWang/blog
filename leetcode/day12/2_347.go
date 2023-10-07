package day12

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type Element struct {
	Num   int
	Count int
}

func byPriority(a, b interface{}) int {
	priorityA := a.(Element).Count
	priorityB := b.(Element).Count
	return -utils.IntComparator(priorityA, priorityB) // "-" descending order
}

func TopKFrequent(nums []int, k int) []int {
	maps := map[int]int{}
	for _, v := range nums {
		maps[v]++
	}
	pq := priorityqueue.NewWith(byPriority)
	for k, v := range maps {
		pq.Enqueue(Element{
			Num:   k,
			Count: v,
		})
	}
	var result []int
	for i := 0; i < k; i++ {
		value, _ := pq.Dequeue()
		result = append(result, value.(Element).Num)
	}
	return result
}
