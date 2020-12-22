package main

import "sort"

func p12933(n int64) int64 {
	var tmp []int
	var ret int64

	for n > 0 {
		tmp = append(tmp, int(n%10))
		n /= 10
	}

	sort.Ints(tmp)
	// sort.Sort(sort.Reverse(sort.IntSlice(tmp)))

	for i := len(tmp) - 1; i >= 0; i-- {
		ret *= 10
		ret += int64(tmp[i])
	}

	return ret
}
