// https://programmers.co.kr/learn/courses/30/lessons/42747

package main

import "sort"

func p42747(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	for i := 0; i < len(citations); i++ {
		if i+1 > citations[i] {
			return i
		}
	}
	return 0
}
