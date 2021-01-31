// https://programmers.co.kr/learn/courses/30/lessons/12939

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func p12939(s string) string {
	tmp := strings.Fields(s)
	p := make([]int, len(tmp))

	for i, t := range tmp {
		p[i], _ = strconv.Atoi(t)
	}
	sort.Ints(p)

	return fmt.Sprint(p[0], p[len(p)-1])
}
