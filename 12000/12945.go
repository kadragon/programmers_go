// https://programmers.co.kr/learn/courses/30/lessons/12945

package main

func p12945(n int) int {
	f := make([]int, 100001)
	f[1] = 1
	f[2] = 1

	for i := 3; i <= n; i++ {
		f[i] = (f[i-1] + f[i-2]) % 1234567
	}

	return f[n]
}
