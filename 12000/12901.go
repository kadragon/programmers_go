package main

func p12901(a int, b int) string {
	p := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
	m := []int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	b += 4

	for i := 1; i < a; i++ {
		b += m[i]
	}

	return p[b%7]
}
