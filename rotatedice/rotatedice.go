package rotatedice

var (
	oppositeFaces = map[int]int{
		1: 6, 6: 1,
		2: 5, 5: 2,
		3: 4, 4: 3,
	}
)

func countTimesAndSide(a []int) (rs map[int]int, side int) {
	min := 0
	for _, v := range a {
		rs[v]++
		min = rs[v]
		side = v
	}
	for k, v := range rs {
		if rs[oppositeFaces[k]] != 0 {
			rs[k] = v * 2
		}
		if rs[k] < min {
			min = rs[k]
			side = k
		}
	}
	return rs, side
}

func Solution(a []int) int {
	for i := 1; i < 7; i++ {

	}
	return 0
}
