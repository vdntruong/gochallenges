package stableperiod

func findMaxCase(a []int, i int) int {
	if i <= len(a)-3 && a[i+1]-a[i] == a[i+2]-a[i+1] {
		return findMaxCase(a, i+1) + 1
	}
	return 0
}

func Solution(a []int) (rs int) {
	for i := 0; i < len(a)-2; i++ {
		if cases := findMaxCase(a, i); cases != 0 {
			rs += cases
		}
	}
	return
}
