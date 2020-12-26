// https://programmers.co.kr/learn/courses/30/lessons/12977

package main

import "sort"

func p12977(nums []int) int {
	answer := 0
	sort.Ints(nums)
	le := len(nums)
	max := nums[le-1] + nums[le-2] + nums[le-3]

	prime := make([]int, max+1)

	for i := 2; i <= max; i++ {
		if prime[i] == 0 {
			prime[i] = 1
			for j := 2; i*j <= max; j++ {
				prime[i*j] = 2
			}
		}
	}

	for i := 0; i < le; i++ {
		for j := i + 1; j < le; j++ {
			for k := j + 1; k < le; k++ {
				if prime[nums[i]+nums[j]+nums[k]] == 1 {
					answer++
				}
			}
		}
	}

	return answer
}
