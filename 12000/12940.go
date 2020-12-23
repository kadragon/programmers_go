// https://programmers.co.kr/learn/courses/30/lessons/12940

package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func p12940(n int, m int) []int {
	var a int
	a = gcd(n, m)

	return []int{a, n / a * m}
}
