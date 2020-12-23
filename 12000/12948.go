// https://programmers.co.kr/learn/courses/30/lessons/12948

package main

import "strings"

func p12948(phoneNumber string) string {
	length := len(phoneNumber)
	ret := strings.Repeat("*", length-4) + phoneNumber[length-4:]

	return ret
}
