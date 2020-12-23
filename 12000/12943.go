// https://programmers.co.kr/learn/courses/30/lessons/12943

package main

func p12943(num int) int {
	var ret int

	for ; num > 1 && ret <= 500; ret++ {
		if num%2 == 0 {
			num /= 2
		} else {
			num *= 3
			num++
		}
	}

	if num != 1 {
		return -1
	}
	return ret
}
