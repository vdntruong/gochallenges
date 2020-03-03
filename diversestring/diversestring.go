package diversestring

import "github.com/sirupsen/logrus"

var (
	A = "a"
	B = "b"
	C = "c"
)

func findMax(center map[string]int) (c string, v int) {
	for ch, vl := range center {
		if vl > v {
			c = ch
			v = vl
		}
	}
	return c, v
}
func findMin(center map[string]int) (c string, v int) {
	c = A
	v = center[A]
	for ch, vl := range center {
		if vl < v {
			c = ch
			v = vl
		}
	}
	return c, v
}
func findMid(center map[string]int) (c string, v int) {
	maxC, _ := findMax(center)
	minC, _ := findMin(center)
	for ch, vl := range center {
		if ch != maxC && ch != minC {
			c = ch
			v = vl
		}
	}
	return c, v
}

func youCanDoIt(min, mid int, maxChar, s string) bool {
	return min == 0 && mid == 0 && string(s[len(s)-1]) == maxChar
}
func makeRs(rs, char string, times int) string {
	for i := 0; i < times; i++ {
		rs += char
	}
	return rs
}

func Solution(a, b, c int) (rs string) {
	center := map[string]int{A: a, B: b, C: c}
	maxC, maxV := findMax(center)
	_, minV := findMin(center)
	midC, midV := findMid(center)

	logrus.Info("maxC", maxC, "maxV", maxV)
	logrus.Info("minV", minV)
	logrus.Info("midC", midC, "midC", midV)
	for youCanDoIt(minV, midV, maxC, rs) {
		if minV != 0 {
			if maxV/minV <= 2 {
				rs = makeRs(rs, maxC, 1)
				center[maxC] -= 1
			} else if maxV/minV > 2 {
				rs = makeRs(rs, maxC, 2)
				center[maxC] -= 2
			}
		} else if maxV <= 2 {
			rs = makeRs(rs, maxC, maxV)
			center[maxC] = 0
		} else {
			rs = makeRs(rs, maxC, 2)
			center[maxC] -= 2
		}

		if minV != 0 {
			if midV/minV <= 2 {
				rs = makeRs(rs, midC, 1)
				center[midC] -= 1
			} else if midV/minV > 2 {
				rs = makeRs(rs, midC, 2)
				center[midC] -= 2
			}
		} else if midV <= 2 {
			rs = makeRs(rs, midC, midV)
			center[midC] = 0
		} else {
			rs = makeRs(rs, midC, 2)
			center[midC] -= 2
		}

		maxC, maxV = findMax(center)
		_, minV = findMin(center)
		midC, midV = findMid(center)
	}

	return rs
}
