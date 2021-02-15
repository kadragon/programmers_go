// https://programmers.co.kr/learn/courses/30/lessons/12904

package main

func p12904(s string) int {
	l := len(s)
	max := 1

	var dp [2500][2500]int

	for i := 0; i < l; i++ {
		dp[i][i] = 1
	}

	for i := 0; i < l-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
		}
	}

	for i := 2; i < l; i++ {
		for j := 0; j < l-i; j++ {
			if dp[j+1][j+i-1] > 0 && s[j] == s[j+i] {
				dp[j][j+i] = dp[j+1][j+i-1] + 2
				if max < dp[j][j+i] {
					max = dp[j][j+i]
				}
			}
		}
	}

	return max
}
