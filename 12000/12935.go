// https://programmers.co.kr/learn/courses/30/lessons/12935

package main

func p12935(arr []int) []int {
	min := 0

	for i := 1; i < len(arr); i++ {
		if arr[min] > arr[i] {
			min = i
		}
	}

	ret := append(arr[0:min], arr[min+1:]...)

	if len(ret) == 0 {
		return []int{-1}
	}

	return ret
}
