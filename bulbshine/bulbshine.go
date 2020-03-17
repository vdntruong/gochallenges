package bulbshine

func Solution(a []int) (rs int) {
	max := 0
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
		if i+1 >= max {
			rs++
		}
	}
	return rs
}

// func OldSolution(A []int) int {
// 	rs := 0
// 	for i := 0; i < len(A); i++ {
// 		fl := true
// 		for j := 0; j <= i; j++ {
// 			if A[j] > i+1 {
// 				fl = false
// 				break
// 			}
// 		}
// 		if fl {
// 			rs++
// 		}
// 	}
// 	return rs
// }

func thisIndexOke(a []int, i, max int) bool {
	if a[i] > max {
		max = a[i]
	}
	if i+1 >= max {
		if i == 0 {
			return true
		}
		return thisIndexOke(a, i-1, max)
	}
	return false
}
func NewSolution(a []int) (rs int) {
	for i := 0; i < len(a); i++ {
		if thisIndexOke(a, i, 0) {
			rs++
		}
	}
	return
}
