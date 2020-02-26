package thering

import "math"

type (
	Point struct {
		x, y int
	}
)

func counted(ps []Point, p Point) bool {
	for _, po := range ps {
		if po == p {
			return true
		}
	}
	return false
}

func Solution(inner, outer int, x []int, y []int) (rs int) {
	if inner < 0 || inner > 9_999 || outer < 1 || outer > 10_000 {
		return rs
	}
	if len(x) != len(y) || inner >= outer {
		return rs
	}
	totalPoint := len(x)
	if totalPoint < 1 ||
		totalPoint > 1_000 {
		return rs
	}

	listPassPoint := make([]Point, 0)
	for i := 0; i < totalPoint; i++ {
		if x[i] < -10_000 || x[i] > 10_000 {
			continue
		}
		if y[i] < -10_000 || y[i] > 10_000 {
			continue
		}
		distance :=
			math.Sqrt(
				math.Pow(float64(x[i]), 2) +
					math.Pow(float64(y[i]), 2))

		if float64(inner) < distance &&
			distance < float64(outer) &&
			!counted(listPassPoint, Point{x[i], y[i]}) {
			listPassPoint = append(listPassPoint, Point{x[i], y[i]})
		}
	}

	return len(listPassPoint)
}
