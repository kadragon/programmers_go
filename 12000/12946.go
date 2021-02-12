// https://programmers.co.kr/learn/courses/30/lessons/12946

package main

var result [][]int

func p12946(n int) [][]int {
	result = [][]int{}
	hanoi(n, 1, 2, 3)
	return result
}

func hanoi(n int, a int, b int, c int) {
	if n > 0 {
		hanoi(n-1, a, c, b)
		result = append(result, []int{a, c})
		hanoi(n-1, b, a, c)
	}
}
