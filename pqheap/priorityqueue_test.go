package pqheap

import (
	"container/heap"
	"fmt"
	"reflect"
	"testing"
)

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func TestPriorityQueue(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
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

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	actual_items := make(map[string]int)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		actual_items[item.value] = item.priority
		//fmt.Printf("%.2d:%s \n", item.priority, item.value)
	}

	expected_items := map[string]int{
		"orange": 5, "pear": 4, "banana": 3, "apple": 2,
	}

	// Output:
	// 05:orange 04:pear 03:banana 02:apple

	eq := reflect.DeepEqual(expected_items, actual_items)
	if eq {
		fmt.Printf("expected: %v\ngot: %v\n", expected_items, actual_items)
	} else {
		t.Fatalf("expected: %v, got: %v\n", expected_items, actual_items)
	}
}
