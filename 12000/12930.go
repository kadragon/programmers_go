package main

func p12930(s string) string {
	ret := make([]rune, len(s))
	var off int

	for i, v := range s {
		ret[i] = v
		if v == ' ' {
			off = i + 1
		} else if (i-off)%2 == 0 {
			if 'a' <= v && v <= 'z' {
				ret[i] = v + 'A' - 'a'
			}
		} else {
			if 'A' <= v && v <= 'Z' {
				ret[i] = v - 'A' + 'a'
			}
		}
	}

	return string(ret)
}
