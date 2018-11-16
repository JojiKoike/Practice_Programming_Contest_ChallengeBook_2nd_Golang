package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string // The value of item
	priority int    // The priority of the item in the queue
	// Index is mandatory for update queue
	// And maintained by the heap.Interface methods.
	index int // The index of the item in the heap
}

// A priorityqueue implements heap.Interface and hold Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We wanna Pop is to give us highest, not lowest priority
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

// Push is to enqueue the value into priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop is to get the value from priority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	// Input Values
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Build Priority Queue
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a New Item and then modify its priority
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Show All Items in priority queue
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Value : %s \t Priority : %d\n", item.value, item.priority)
	}

}
