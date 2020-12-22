package main

import "sort"

func p12910(arr []int, divisor int) []int {
	var ans []int
	for _, p := range arr {
		if p%divisor == 0 {
			ans = append(ans, p)
		}
	}
	sort.Ints(ans)

	if len(ans) > 0 {
		return ans
	}

	return []int{-1}
}
