// https://programmers.co.kr/learn/courses/30/lessons/42746

package main

import (
	"sort"
	"strconv"
	"strings"
)

func p42746(numbers []int) string {
	var str []string
	for _, v := range numbers {
		str = append(str, strconv.Itoa(v))
	}

	var li, lj int
	var a, b byte

	sort.Slice(str, func(i, j int) bool {
		li = len(str[i])
		lj = len(str[j])

		for k := 0; k < li || k < lj; k++ {
			if k < li {
				a = str[i][k]
			} else {
				a = str[i][0]
			}
			if k < lj {
				b = str[j][k]
			} else {
				b = str[j][0]
			}

			if a == b {
				continue
			}

			return a > b
		}

		return str[i][len(str[i])-1] > str[j][len(str[j])-1]
	})

	ans := strings.Join(str, "")

	if ans[0] == '0' {
		return "0"
	}

	return ans
}
