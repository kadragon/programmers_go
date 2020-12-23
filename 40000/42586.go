// https://programmers.co.kr/learn/courses/30/lessons/42586

package main

func p42586(progresses []int, speeds []int) []int {
	var ret []int
	var t, time int

	for i, v := range progresses {
		if v+time*speeds[i] < 100 {
			ret = append(ret, t)

			time = (100 - v) / speeds[i]
			if (100-v)%speeds[i] > 0 {
				time++
			}

			t = 1
		} else {
			t++
		}
	}

	ret = append(ret, t)

	return ret[1:]
}
