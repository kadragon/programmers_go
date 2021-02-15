// https://programmers.co.kr/learn/courses/30/lessons/49191

package main

import (
	"sort"
)

var d [101][101]bool
var w []int
var p [][]int

func p49191(n int, results [][]int) int {
	w = make([]int, n+1)
	p = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = []int{}
		w[i] = 1
	}

	for _, r := range results {
		p[r[0]] = append(p[r[0]], r[1])
	}

	for i := 1; i <= n; i++ {
		wIn(i, i)
	}

	sort.Ints(w)
	result := 0

	ch := make([]bool, 101)

	for i := n; i > 0; i-- {
		if w[i] == i && !ch[i] {
			result++
		} else {
			ch[w[i]] = true
		}
	}

	return result
}

func wIn(s int, now int) {
	for _, r := range p[now] {
		if !d[r][s] {
			d[r][s] = true
			w[r]++
			wIn(s, r)
		}
	}
}
