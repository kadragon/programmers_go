// https://programmers.co.kr/learn/courses/30/lessons/43162

package main

func p43162_2(n int, computers [][]int) int {
	visited := make([]bool, n)
	result := 0

	for i := 0; i < n; i++ {
		if visited[i] == false {
			result++
			dfs(n, computers, visited, i)
		}
	}

	return result
}

func dfs(n int, computers [][]int, visited []bool, now int) {
	visited[now] = true

	for i := 0; i < n; i++ {
		if computers[now][i] == 1 && visited[i] == false {
			dfs(n, computers, visited, i)
		}
	}
}
