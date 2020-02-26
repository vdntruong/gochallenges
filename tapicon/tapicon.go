package tapicon

import "math"

type (
	Point struct {
		x, y int
	}
)

const RADIUS = 20

func (p Point) isInCircularArea(otherPoint Point) (rs bool) {
	distance := math.Sqrt(
		math.Pow(float64(p.x)-float64(otherPoint.x), 2) +
			math.Pow(float64(p.y)-float64(otherPoint.y), 2))
	return distance <= float64(RADIUS)
}

func Solution(A []int, B []int, X int, Y int) int {
	if len(A) != len(B) {
		return -1
	}
	for i := 0; i < len(A); i++ {
		if (Point{X, Y}).isInCircularArea(Point{A[i], B[i]}) {
			return i
		}
	}
	return -1
}
