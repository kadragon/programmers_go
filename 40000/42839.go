// https://programmers.co.kr/learn/courses/30/lessons/42839

package main

import (
	"strconv"
)

var result int
var checked map[int]bool
var byteString []rune

func p42839(numbers string) int {
	checked = make(map[int]bool)
	checked[0], checked[1] = false, false
	result = 0

	byteString = []rune{}
	for _, c := range numbers {
		byteString = append(byteString, c)
	}
	for i := 1; i <= len(numbers); i++ {
		permutation(len(numbers), i, 0)
	}

	return result
}

// 순열(재귀)
// https://iyk2h.tistory.com/80
func permutation(n, r, depth int) {
	if r == depth {
		resultCheck(string(byteString[:depth]))
		return
	}
	for i := depth; i < n; i++ {
		byteString[i], byteString[depth] = byteString[depth], byteString[i]
		permutation(n, r, depth+1)
		byteString[i], byteString[depth] = byteString[depth], byteString[i]
	}
}

func resultCheck(number string) {
	target, _ := strconv.Atoi(number)
	if _, ok := checked[target]; !ok {
		checked[target] = checkPrime(target)
		if checked[target] {
			result++
		}
	}
}

func checkPrime(number int) bool {
	if number == 2 {
		return true
	}

	if number%2 == 0 {
		return false
	}

	for i := 3; i*i <= number; i = i + 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}
