// https://programmers.co.kr/learn/courses/30/lessons/42587

package main

func p42587(priorities []int, location int) int {
	var idx, cnt, tmp, h int
	length := len(priorities)
	for i := 9; i > 0; i-- {
		for j := 0; j < length; j++ {
			h = (idx + j) % length
			if priorities[h] == i {
				tmp = h
				cnt++
				if h == location {
					return cnt
				}
			}
		}
		idx = tmp
	}

	return 0
}
