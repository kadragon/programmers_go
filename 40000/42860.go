// https://programmers.co.kr/learn/courses/30/lessons/42860

package main

func p42860(name string) int {
	m := make([]int, len(name))

	for i, v := range name {
		m[i] = int(v - 65)
		if m[i] > 13 {
			m[i] = 26 - m[i]
		}
	}

	var ans, max, length, idx int
	for i, v := range m {
		if v == 0 {
			length++
			if max < length {
				max = length
				idx = i
			}
		} else {
			ans += v
			length = 0
		}
	}

	a := idx - max

	if a < max {
		ans -= max
		ans += a
	}

	return ans + len(name) - 1
}
