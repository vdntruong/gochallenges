package diversestring

import "math"

func stopNow(a, b, c int) bool {
	return (a == 0 && b == 0) ||
		(a == 0 && c == 0) ||
		(b == 0 && c == 0)
}

func addStr(str, s string, times int) string {
	for i := 0; i < times; i++ {
		str += s
	}
	return str
}
func findOrder(a, b, c int) (max, mid, min int) {
	max = int(math.Max(math.Max(float64(a), float64(b)), float64(c)))
	min = int(math.Min(math.Min(float64(a), float64(b)), float64(c)))
	if max >= a && a >= min {
		mid = a
		return
	}
	if max >= b && b >= min {
		mid = b
		return
	}
	mid = c
	return
}
func findRepeatTimes(a int) int {
	if a <= 2 {
		return 1
	}
	return 2
}

func findMaxMidChar(disCenter map[string]int, max, min int) (maxC, midC, minC string) {
	for char, count := range disCenter {
		if count == max {
			maxC = char
		} else if count == min {
			minC = char
		} else {
			midC = char
		}
	}
	return
}

func makeRs(disCenter map[string]int) (rs string) {
	for {
		if stopNow(disCenter["a"], disCenter["b"], disCenter["c"]) {
			max, _, min := findOrder(disCenter["a"], disCenter["b"], disCenter["c"])
			maxC, _, _ := findMaxMidChar(disCenter, max, min)
			if max <= 2 {
				rs = addStr(rs, maxC, max)
			} else {
				rs = addStr(rs, maxC, 2)
			}
			return rs
		}
		max, mid, min := findOrder(disCenter["a"], disCenter["b"], disCenter["c"])
		maxC, midC, _ := findMaxMidChar(disCenter, max, min)

		divMaxMid := 0
		if mid != 0 {
			divMaxMid = max / mid
		} else if mid <= 2 {
			divMaxMid = mid
		} else {
			divMaxMid = 2
		}

		divMaxMin := 0
		if min != 0 {
			divMaxMin = max / min
		} else if max <= 2 {
			divMaxMin = max
		} else {
			divMaxMin = 2
		}
		maxTimes := findRepeatTimes(divMaxMin)
		midTimes := findRepeatTimes(divMaxMid)

		rs = addStr(rs, maxC, maxTimes)
		disCenter[maxC] -= maxTimes
		rs = addStr(rs, midC, midTimes)
		disCenter[midC] -= midTimes
	}
}

func Solution(a, b, c int) (rs string) {
	distributeCenter := map[string]int{"a": a, "b": b, "c": c}
	rs = makeRs(distributeCenter)
	return rs
}
