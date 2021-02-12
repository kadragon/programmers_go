// https://programmers.co.kr/learn/courses/30/lessons/12914

package main

func p12914(n int) int64 {
	a, b := 1, 1

	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%1234567
	}

	return int64(b)
}
