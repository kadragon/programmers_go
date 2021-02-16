// https://programmers.co.kr/learn/courses/30/lessons/43165

package main

func p43165(numbers []int, target int) int {
	result := 0
	s(0, 0, target, numbers, &result)

	return result
}

func s(depth, now, target int, numbers []int, result *int) {
	if depth == len(numbers) {
		if now == target {
			*result++
		}

		return
	}

	s(depth+1, now+numbers[depth], target, numbers, result)
	s(depth+1, now-numbers[depth], target, numbers, result)
}

// Other Solution
func sol(numbers []int, target int) int {
	if len(numbers) == 0 && target == 0 {
		return 1
	} else if len(numbers) == 0 {
		return 0
	}
	return sol(numbers[1:], target-numbers[0]) + sol(numbers[1:], target+numbers[0])
}
