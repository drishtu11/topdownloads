package pqheap

import (
	"container/heap"
	"fmt"
)

// A PriorityQueue implements heap.Interface and holds Items.
// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// findTopKDownloads : finds the top K downloads of files from given repo
func FindTopKDownloads(countMap map[string]int, num int) {
	pq := make(PriorityQueue, len(countMap))
	i := 0
	for value, priority := range countMap {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Take the items out; they arrive in decreasing priority order.
	count := 0
	fmt.Printf("----------------------------------------\n")
	fmt.Printf("Top %d Downloads\n", num)
	fmt.Printf("----------------------------------------\n")
	for pq.Len() > 0 && count < num {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Artifact : %s\nDownloads : %d\n\n", item.value, item.priority)
		count++
	}
	fmt.Printf("----------------------------------------\n")
	fmt.Println("")
}
