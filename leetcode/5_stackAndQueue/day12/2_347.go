package day12

import (
	"fmt"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)
type Element struct {
	Name     int
	Priority int
}

func byPriority(a, b interface{}) int {
	priorityA := a.(Element).Priority
	priorityB := b.(Element).Priority
	return -utils.IntComparator(priorityA, priorityB) // "-" descending order
}

func topKFrequent(nums []int, k int) []int {
	maps := map[int]int{}
	for _, v := range nums {
		maps[v]++
	}
	pq := priorityqueue.NewWith(byPriority)
	for k, v := range maps {
		pq.Enqueue(Element{
			Name:     k,
			Priority: v,
		})
	}
	fmt.Println(pq)
	var result []int
	for i := 0; i < k; i++ {
		value, _ := pq.Dequeue()
		fmt.Println(value)
		result = append(result, value.(Element).Name)
	}
	return result
}