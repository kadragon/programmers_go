// https://programmers.co.kr/learn/courses/30/lessons/42842

package main

func p42842(brown int, yellow int) []int {
	c := (brown - 4) / 2

	for b := 1; b < c; b++ {
		if b*(c-b) == yellow {
			return []int{c - b + 2, b + 2}
		}
	}
	return []int{}
}
