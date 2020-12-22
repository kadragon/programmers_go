package main

func p12932(n int64) []int {
	var ret []int

	for n > 0 {
		ret = append(ret, int(n%10))
		n /= 10
	}

	return ret
}
