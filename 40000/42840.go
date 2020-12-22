//https://programmers.co.kr/learn/courses/30/lessons/42840

package main

func p42840(answers []int) []int {
	order := [][]int{
		{1, 2, 3, 4, 5},
		{2, 1, 2, 3, 2, 4, 2, 5},
		{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}}
	length := []int{5, 8, 10}

	a := make([]int, 3)
	var ans []int
	max := 0

	for i := 0; i < len(answers); i++ {
		for j := 0; j < 3; j++ {
			if answers[i] == order[j][i%length[j]] {
				a[j]++
			}
		}
	}

	for j := 0; j < 3; j++ {
		if max < a[j] {
			max = a[j]
		}
	}

	for j := 0; j < 3; j++ {
		if max == a[j] {
			ans = append(ans, j+1)
		}
	}

	return ans
}
