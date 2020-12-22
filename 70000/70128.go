package main

func p70128(a []int, b []int) int {
	var ret int
	for i := 0; i < len(a); i++ {
		ret += a[i] * b[i]
	}
	return ret
}
