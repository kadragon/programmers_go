// https://programmers.co.kr/learn/courses/30/lessons/12969

package main

import (
	"fmt"
	"strings"
)

func p12969() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 0; i < b; i++ {
		fmt.Println(strings.Repeat("*", a))
	}
}
