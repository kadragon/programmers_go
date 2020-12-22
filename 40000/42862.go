package main

func p42862(n int, lost []int, reserve []int) int {
	check := make([]bool, n+2)
	fList := make([]int, 0, len(lost))
	for _, r := range reserve {
		check[r] = true
	}

	for _, l := range lost {
		if check[l] {
			check[l] = false
		} else {
			fList = append(fList, l)
		}
	}

	for _, l := range fList {
		if check[l-1] {
			check[l-1] = false
		} else if check[l+1] {
			check[l+1] = false
		} else {
			n--
		}
	}

	return n
}
