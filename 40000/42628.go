// https://programmers.co.kr/learn/courses/30/lessons/42628

package main

import (
	"container/heap"
	"strconv"
	"strings"
)

// IntHeap ...
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push ...
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop ...
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func p42628(operations []string) []int {
	maxpq := &IntHeap{}
	minpq := &IntHeap{}
	count := 0

	heap.Init(maxpq)
	heap.Init(minpq)

	for _, oper := range operations {
		t := strings.Split(oper, " ")
		p, _ := strconv.Atoi(t[1])

		if t[0] == "I" {
			heap.Push(maxpq, -p)
			heap.Push(minpq, p)
			count++
		} else if t[0] == "D" && count > 0 {
			if p == 1 {
				heap.Pop(maxpq)
			} else {
				heap.Pop(minpq)
			}
			count--
		}

		if count == 0 {
			for maxpq.Len() > 0 {
				heap.Pop(maxpq)
			}
			for minpq.Len() > 0 {
				heap.Pop(minpq)
			}
		}
	}

	if count == 0 {
		return []int{0, 0}
	}
	a := heap.Pop(minpq)
	b := heap.Pop(maxpq)
	return []int{-b.(int), a.(int)}
}
