package collectingapples

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type (
	SomeThing struct {
		forAlice []int
		forBob   []int
	}
)

func (s *SomeThing) CountTotal() int {
	return totalA(s.forAlice) + totalA(s.forBob)
}

// util function
func totalA(A []int) (total int) {
	for _, v := range A {
		total += v
	}
	return
}
func totalSomeOne(len, n int) int {
	return len - n + 1
}
func maxSomeThing(lenA, K, L int) int {
	return totalSomeOne(lenA, K) + totalSomeOne(lenA-K, L)
}

func Solution(A []int, K int, L int) int {
	if K+L > len(A) {
		return -1
	}
	if K+L == len(A) {
		return totalA(A)
	}

	totalMax := 0
	bagChannel := make(chan SomeThing, maxSomeThing(len(A), K, L))
	wg := sync.WaitGroup{}

	go func() {
		
		stepK := 0
		for {
			forAlice := A[:K]
			leftA := A[:stepK]
			rightA := A[stepK:]

			if len(leftA) >= L {

			}
			if len(rightA) >= L {

			}
			bagChannel<- SomeThing{forAlice: forAlice, forBob: }
			stepK++
		}
	}()

	go func() {
		wg.Add(1)
		for st := range bagChannel {
			logrus.Info("Total is: " + string(st.CountTotal()))
		}
		wg.Done()
	}()

	wg.Wait()
	return totalMax
}
