package cuttingtree

func stillOk(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func Solution(a []int) (rs int) {
	for i := 0; i < len(a)-1; i++ {
		temp := make([]int, len(a))
		copy(temp, a[:])
		if stillOk(append(temp[:i], temp[i+1:]...)) {
			rs++
		}
	}
	if stillOk(a[:len(a)-1]) {
		rs++
	}
	return rs
}
