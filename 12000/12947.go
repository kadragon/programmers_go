// https://programmers.co.kr/learn/courses/30/lessons/12947

package main

func p12947(x int) bool {
	d := 0
	t := x

	for t > 0 {
		d += t % 10
		t /= 10
	}

	return x%d == 0
}
