// https://programmers.co.kr/learn/courses/30/lessons/12944

package main

func p12944(arr []int) float64 {
	var ret float64
	for _, a := range arr {
		ret += float64(a)
	}

	return ret / float64(len(arr))
}
