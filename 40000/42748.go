package main

import "sort"

func p42784(array []int, commands [][]int) []int {
	var ans []int
	for _, command := range commands {
		tmp := append([]int{}, array[command[0]-1:command[1]]...)
		sort.Ints(tmp)
		ans = append(ans, tmp[command[2]-1])
	}

	return ans
}
