package main

func p12912(a int, b int) int64 {
	if a > b {
		var t int
		t = a
		a = b
		b = t
	}
	A := int64(a)
	B := int64(b)

	return (A + B) * (B - A + 1) / 2
}
