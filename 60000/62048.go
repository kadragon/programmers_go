// https://programmers.co.kr/learn/courses/30/lessons/62048

package main

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func p62048(w int, h int) int64 {
	if w > h {
		w, h = h, w
	}

	return int64(w)*int64(h) - int64(w+h-gcd(w, h))
}
