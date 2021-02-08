// https://programmers.co.kr/learn/courses/30/lessons/12949

package main

func p12949(arr1 [][]int, arr2 [][]int) [][]int {
	a := len(arr1)
	b := len(arr2[0])
	c := len(arr1[0])

	result := make([][]int, a)
	for i := 0; i < a; i++ {
		result[i] = make([]int, b)
	}

	for k := 0; k < c; k++ {
		for i := 0; i < a; i++ {
			for j := 0; j < b; j++ {
				result[i][j] += arr1[i][k] * arr2[k][j]
			}
		}
	}

	return result
}
