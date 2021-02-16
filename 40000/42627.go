// https://programmers.co.kr/learn/courses/30/lessons/42627

package main

import (
	"container/heap"
	"sort"
)

// Item ...
type Item struct {
	value int
	time  int
}

// PQ ...
type PQ []*Item

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}

// Swap ...
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push ...
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

// Pop ...
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

// Top ...
func (pq *PQ) Top() interface{} {
	old := *pq
	return old[len(old)-1]
}

func p42627(jobs [][]int) int {
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i][0] < jobs[j][0] {
			return true
		} else if jobs[i][0] == jobs[j][0] {
			return jobs[i][1] < jobs[j][1]
		}
		return false
	})

	pq := PQ{}
	heap.Init(&pq)

	now, wait := 0, 0

	heap.Push(&pq, &Item{jobs[0][0], jobs[0][1]})
	i := 1

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if now < item.value {
			now = item.value
		}

		now += item.time
		wait += now - item.value

		for ; i < len(jobs); i++ {
			if jobs[i][0] <= now {
				heap.Push(&pq, &Item{jobs[i][0], jobs[i][1]})
			} else {
				break
			}
		}

		if pq.Len() == 0 && i < len(jobs) {
			heap.Push(&pq, &Item{jobs[i][0], jobs[i][1]})
			i++
		}
	}

	return wait / len(jobs)
}
