package main

import "sort"

func p12917(s string) string {
	a := []byte(s)
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	return string(a)
}
