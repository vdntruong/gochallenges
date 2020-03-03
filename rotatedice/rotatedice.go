package rotatedice

var (
	oppositeFaces = map[int]int{
		1: 6, 6: 1,
		2: 5, 5: 2,
		3: 4, 4: 3,
	}
)

func findMin(a []int) int {
	min := a[0]
	for _, v := range a {
		if min > v {
			min = v
		}
	}
	return min
}
func Solution(a []int) int {

	siteCenter := []int{}
	for i := 0; i < len(a); i++ {
		count := 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				continue
			}
			if oppositeFaces[a[i]] == a[j] {
				count++
			}
			count++
		}
		siteCenter = append(siteCenter, count)
	}

	return findMin(siteCenter)
}
