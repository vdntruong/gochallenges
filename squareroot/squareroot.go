package squareroot

import "math"

func countTimes(number int) (rs int) {
	sq := math.Sqrt(float64(number))
	if sq == float64(int64(sq)) {
		return 1 + countTimes(int(sq))
	}
	return rs
}

func Solution(a, b int) (rs int) {
	for ; a <= b; a++ {
		maxTimes := countTimes(a)
		if maxTimes > rs {
			rs = maxTimes
		}
	}
	return rs
}
