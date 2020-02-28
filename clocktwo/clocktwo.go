package clocktwo

import (
	"math/rand"
	"sync"
)

const (
	MaxHour   = 23
	MaxMinute = 59
)

type (
	SetCase [][4]int

	Tim struct {
		hour, minute int
	}
)

func Solution(mainAr []int) int {
	var wg sync.WaitGroup

	rs := 0
	length := len(mainAr)

	totalCase := factorial(length) / factorial(length-4)
	ourIndexSet := make([][4]int, 0)
	outSetChannel := make(chan []int)

	// arrangement index
	wg.Add(1)
	go func() {
		defer wg.Done()

		for len(ourIndexSet) < totalCase {
			tempAr := make([]int, length)
			copy(tempAr, mainAr)
			var found []int

			for {
				i := rand.Intn(length)
				if tempAr[i] != -1 {
					found = append(found, i)
					if len(found) == 4 {
						break
					}
					tempAr[i] = -1
				}
			}

			newAr := [4]int{found[0], found[1], found[2], found[3]}
			if notContain(newAr, ourIndexSet) {
				ourIndexSet = append(ourIndexSet, newAr)
				outSetChannel <- found
			}
		}

		close(outSetChannel)
	}()

	// check valid
	wg.Add(1)
	go func() {
		defer wg.Done()
		ourTimes := make([]Tim, 0)

		for newCase := range outSetChannel {
			newTim := createTim(newCase)
			if checkTime(newTim) && notDuplicateTime(ourTimes, newTim) {
				ourTimes = append(ourTimes, newTim)
				rs++
			}
		}
	}()

	wg.Wait()
	return rs
}

func createTim(outTime []int) Tim {
	return Tim{
		hour:   createDecimalNumber(outTime[0], outTime[1]),
		minute: createDecimalNumber(outTime[2], outTime[3]),
	}
}

func checkTime(outTime Tim) (rs bool) {
	return outTime.hour <= MaxHour && outTime.minute <= MaxMinute
}

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

func factorial(i int) int {
	if i > 1 {
		return i * factorial(i-1)
	}
	return 1
}
