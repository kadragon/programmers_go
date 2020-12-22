package main

func p12918(s string) bool {
	t := []byte(s)

	for _, i := range t {
		if i < 48 || i > 57 {
			return false
		}
	}

	if len(t) == 4 || len(t) == 6 {
		return true
	}

	return false
}
