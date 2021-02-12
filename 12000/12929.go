// https://programmers.co.kr/learn/courses/30/lessons/12929

package main

func p12929(n int) int {
	a, b := n, n
	return next(a-1, b, 1)
}

func next(a, b, r int) int {
	if a < 0 || b < 0 || r < 0 {
		return 0
	}
	if a == 0 && b == 0 {
		return 1
	}
	return next(a-1, b, r+1) + next(a, b-1, r-1)
}
