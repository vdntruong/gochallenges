package squareroot

import (
	"math"
)

func youCanDoIt(number int) int {
	sq := math.Sqrt(float64(number))
	if sq == float64(int64(sq)) {
		return 1 + youCanDoIt(int(sq))
	}
	return 0
}

func Solution(a, b int) (rs int) {
	for ; a <= b; a++ {
		curMax := youCanDoIt(a)
		if curMax > rs {
			rs = curMax
		}
	}
	return rs
}
