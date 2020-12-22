package main

func p12926(s string, n int) string {
	ret := make([]rune, len(s))
	off := rune(n)

	for i, v := range s {
		if 'A' <= v && v <= 'Z' {
			ret[i] = (v-'A'+off)%26 + 'A'
		} else if 'a' <= v && v <= 'z' {
			ret[i] = (v-'a'+off)%26 + 'a'
		} else {
			ret[i] = v
		}
	}

	return string(ret)
}
