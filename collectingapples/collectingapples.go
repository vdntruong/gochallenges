package collectingapples

func Solution(a []int, k, l int) int {
	max := l
	min := k
	if k == whoMax(k, l) {
		max = k
		min = l
	}
	curMax := -1

	for i := 0; i < len(a)-max+1; i++ {
		ofMax := a[i : i+max]
		pre := a[:i]
		post := a[i+max:]
		if len(pre) < min && len(post) < min {
			continue
		}

		curMaxOfK := -1
		if countWeight(ofMax) > curMaxOfK {
			curMaxOfK = countWeight(ofMax)
		}

		curMaxOfL := -1
		if len(pre) >= min {
			if findMaxWithRange(pre, min) > curMaxOfL {
				curMaxOfL = findMaxWithRange(pre, min)
			}
		}
		if len(post) >= min {
			if findMaxWithRange(post, min) > curMaxOfL {
				curMaxOfL = findMaxWithRange(post, min)
			}
		}

		if curMaxOfK != -1 && curMaxOfL != -1 && curMaxOfK+curMaxOfL > curMax {
			curMax = curMaxOfK + curMaxOfL
		}
	}
	return curMax
}

func findMaxWithRange(a []int, r int) int {
	curMax := 0
	for j := 0; j < len(a)-r+1; j++ {
		if countWeight(a[j:j+r]) > curMax {
			curMax = countWeight(a[j : j+r])
		}
	}
	return curMax
}
func whoMax(k, l int) int {
	if k > l {
		return k
	}
	return l
}
func countWeight(a []int) (rs int) {
	for _, v := range a {
		rs += v
	}
	return rs
}

// OLD

// type (
// 	SomeThing struct {
// 		forAlice []int
// 		forBob   []int
// 	}
// )

// func (s *SomeThing) CountTotal() int {
// 	return totalA(s.forAlice) + totalA(s.forBob)
// }

// // util function
// func totalA(A []int) (total int) {
// 	for _, v := range A {
// 		total += v
// 	}
// 	return
// }

// func totalSomeOne(len, n int) int {
// 	return len - n + 1
// }

// func maxSomeThing(lenA, K, L int) int {
// 	return totalSomeOne(lenA, K) + totalSomeOne(lenA-K, L)
// }

// func Solution(A []int, K int, L int) int {
// 	if K+L > len(A) {
// 		return -1
// 	}
// 	if K+L == len(A) {
// 		return totalA(A)
// 	}
// 	totalMax := 0
// 	bagChannel := make(chan SomeThing, maxSomeThing(len(A), K, L))
// 	wg := sync.WaitGroup{}

// 	// go func() {
// 	// 	i := 0
// 	// 	forAlice := A[i:K]
// 	// 	maxForAlice := totalA(forAlice)
// 	// 	for _; i < totalSomeOne(len(A), K)-1; i++ {
// 	// 		if totalA(A[i:K+i]) > maxForAlice {
// 	// 			maxForAlice := totalA(A[i : K+i])
// 	// 		}
// 	// 	}
// 	// 	for {
// 	// 		leftA := A[:stepK]
// 	// 		rightA := A[stepK:]

// 	// 		if len(leftA) >= L {

// 	// 		}
// 	// 		if len(rightA) >= L {

// 	// 		}
// 	// 		// bagChannel <- SomeThing{forAlice: forAlice, forBob: }
// 	// 		stepK++
// 	// 	}
// 	// }()

// 	go func() {
// 		wg.Add(1)
// 		for st := range bagChannel {
// 			logrus.Info("Total is: " + string(st.CountTotal()))
// 		}
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	return totalMax
// }
