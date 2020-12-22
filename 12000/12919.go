package main

import "fmt"

func p12919(seoul []string) string {
	for i, s := range seoul {
		if s == "Kim" {
			return fmt.Sprintf("김서방은 %d에 있다", i)
		}
	}

	return ""
}
