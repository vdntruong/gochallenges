package collectingapples

func Solution(a []int, k, l int) (rs int) {
	for i := 0; i < len(a)-k+1; i++ {
		ofMax := a[i : i+k]
		pre := a[:i]
		post := a[i+k:]

		if len(pre) < l && len(post) < l {
			continue
		}
		curMaxOfK := countWeight(ofMax)
		curMaxOfL := findMaxInTwoSlice(pre, post, l)

		if curMaxOfK+curMaxOfL > rs {
			rs = curMaxOfK + curMaxOfL
		}
	}
	return rs
}
func findMaxInTwoSlice(a, b []int, l int) (rs int) {
	if len(a) >= l {
		maxSlideA := findMaxWithRange(a, l)
		if maxSlideA > rs {
			rs = maxSlideA
		}
	}
	if len(b) >= l {
		maxSliceB := findMaxWithRange(b, l)
		if maxSliceB > rs {
			rs = maxSliceB
		}
	}
	return rs
}
func findMaxWithRange(a []int, r int) (rs int) {
	for j := 0; j < len(a)-r+1; j++ {
		if countWeight(a[j:j+r]) > rs {
			rs = countWeight(a[j : j+r])
		}
	}
	return rs
}
func countWeight(a []int) (rs int) {
	for _, v := range a {
		rs += v
	}
	return rs
}
