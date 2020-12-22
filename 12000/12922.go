package main

func p12922(n int) string {
	t := ""
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			t += "수"
		} else {
			t += "박"
		}
	}
	return t
}
