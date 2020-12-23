// https://programmers.co.kr/learn/courses/30/lessons/12950

package main

func p12950(arr1 [][]int, arr2 [][]int) [][]int {
	for i := range arr1 {
		for j := range arr1[i] {
			arr1[i][j] += arr2[i][j]
		}
	}

	return arr1
}
