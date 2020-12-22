package main

func p12928(n int) int {
	ret := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ret += i
			ret += n / i
		}
		if i*i == n {
			ret -= i
		}
	}

	return ret
}
