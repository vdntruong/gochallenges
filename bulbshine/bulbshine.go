package bulbshine

func Solution(a []int) (rs int) {
	curMax := 0
	for i := 0; i < len(a); i++ {
		if a[i] > curMax {
			curMax = a[i]
		}
		if curMax <= i+1 {
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
