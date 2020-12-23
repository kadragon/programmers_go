// https://programmers.co.kr/learn/courses/30/lessons/12954

package main

func p12954(x int, n int) []int64 {
	var ret []int64

	for i := 1; i <= n; i++ {
		ret = append(ret, int64(x)*int64(i))
	}

	return ret
}
