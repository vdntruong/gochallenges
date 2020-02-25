package clock

import (
	"math/rand"
	"sync"
)

const (
	MinHour = 0
	MaxHour = 23

	MinMinute = 0
	MaxMinute = 59

	TotalCase = 24
)

type (
	SetCase [][4]int
	Tim     struct {
		hour, minute int
	}
)

func createDecimalNumber(a, b int) int {
	return a*10 + b
}

func notContain(n [4]int, s SetCase) bool {
	for _, ar := range s {
		if n == ar {
			return false
		}
	}
	return true
}

func notDuplicateTime(table []Tim, t Tim) bool {
	for _, ta := range table {
		if ta.hour == t.hour && ta.minute == t.minute {
			return false
		}
	}
	return true
}

func Solution(a, b, c, d int) int {
	if a < 0 || b < 0 || c < 0 || d < 0 {
		return 0
	}
	if a > 9 || b > 9 || c > 9 || d > 9 {
		return 0
	}
	rs := 0

	mainAr := [4]int{a, b, c, d}
	var wg sync.WaitGroup

	ourIndexSet := make([][4]int, 0)
	outSetChannel := make(chan []int)

	// arrangement index
	wg.Add(1)
	go func() {
		for {
			if len(ourIndexSet) == 24 {
				break
			}
			tempAr := mainAr
			var found []int
			for {
				i := rand.Intn(4)
				if tempAr[i] != -1 {
					found = append(found, i)
					if len(found) == 4 {
						break
					}
					tempAr[i] = -1
				}
			}
			outSetChannel <- found
		}
		close(outSetChannel)
		wg.Done()
	}()

	// push to out index set
	wg.Add(1)
	go func() {
		for {
			select {
			case newCase, ok := <-outSetChannel:
				if !ok {
					wg.Done()
					return
				}
				newAr := [4]int{newCase[0], newCase[1], newCase[2], newCase[3]}
				if notContain(newAr, ourIndexSet) {
					ourIndexSet = append(ourIndexSet, newAr)
				}
			}
		}
	}()

	// check valid
	wg.Add(1)
	go func() {
		for {
			if len(ourIndexSet) == 24 {
				ourTimes := make([]Tim, 0)
				for _, ar := range ourIndexSet {
					hour := createDecimalNumber(mainAr[ar[0]], mainAr[ar[1]])
					minute := createDecimalNumber(mainAr[ar[2]], mainAr[ar[3]])
					if hour < 24 && minute < 60 {
						newTim := Tim{hour: hour, minute: minute}
						if notDuplicateTime(ourTimes, newTim) {
							rs++
							ourTimes = append(ourTimes, newTim)
						}
					}
				}
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()
	return rs
}

// func ValueSolution(a, b, c, d int) int {
// 	l := [4]int{a, b, c, d}
// 	rs := 0
// 	var wg sync.WaitGroup

// 	ourSet := make([][4]int, 0)
// 	pushToSet := make(chan []int)

// 	wg.Add(1)
// 	go func(p chan []int, w *sync.WaitGroup) {
// 		for {
// 			if len(ourSet) == 24 {
// 				break
// 			}
// 			tempL := l
// 			var found []int
// 			for {
// 				j := rand.Intn(4)
// 				if tempL[j] != -1 {
// 					found = append(found, tempL[j])
// 					if len(found) == 4 {
// 						break
// 					}
// 					tempL[j] = -1
// 				}
// 			}
// 			p <- found
// 		}

// 		close(p)
// 		w.Done()
// 	}(pushToSet, &wg)

// 	wg.Add(1)
// 	go func(pu <-chan []int) {
// 		for {
// 			select {
// 			case newCase, ok := <-pu:
// 				if !ok {
// 					wg.Done()
// 					return
// 				}
// 				newAr := [4]int{newCase[0], newCase[1], newCase[2], newCase[3]}
// 				if notContain(newAr, ourSet) {
// 					ourSet = append(ourSet, newAr)
// 				}
// 			}
// 		}
// 	}(pushToSet)

// 	wg.Add(1)
// 	go func() {
// 		for {
// 			if len(ourSet) == 24 {
// 				for _, ar := range ourSet {
// 					if createDecimalNumber(ar[0], ar[1]) < 24 && createDecimalNumber(ar[2], ar[3]) < 60 {
// 						rs++
// 					}
// 				}
// 				break
// 			}
// 		}
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	return rs
// }
