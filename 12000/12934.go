import "math"

func p12934(n int64) int64 {
	t := int64(math.Sqrt(float64(n)))

	if n == t*t {
		return (t + 1) * (t + 1)
	}

	return -1
}
