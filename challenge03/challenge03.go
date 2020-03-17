package challenge03

func Solution(ar []int) (rs int) {
	if validateTrees(ar) {
		return rs
	}

	for i := 0; i < len(ar); i++ {
		temp := make([]int, len(ar))
		copy(temp, ar)
		if validateTrees(append(temp[:i], temp[i+1:]...)) {
			rs++
		}
	}

	if rs == 0 {
		return -1
	}
	return rs
}

func validateTrees(a []int) bool {
	asc := a[0] < a[1]
	for i := 0; i < len(a)-1; i++ {
		if asc && a[i] >= a[i+1] {
			return false
		}
		if !asc && a[i] <= a[i+1] {
			return false
		}
		asc = !asc
	}
	return true
}
