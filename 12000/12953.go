// https://programmers.co.kr/learn/courses/30/lessons/12953

package main

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func p12953(arr []int) int {
	ans := arr[0]
	for i := 1; i < len(arr); i++ {
		ans = lcm(ans, arr[i])
	}

	return ans
}
