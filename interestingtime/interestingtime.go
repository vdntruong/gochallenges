package interestingtime

import (
	"math"
	"strconv"
	"strings"
)

type (
	Time struct {
		hour, minute, second int
	}
)

func decimalOne(a int) int {
	return a / 10
}
func decimalTwo(a int) int {
	return a % 10
}

func findTimeParts(str string) (h, m, s int) {
	arTime := strings.Split(str, ":")
	h, _ = strconv.Atoi(arTime[0])
	m, _ = strconv.Atoi(arTime[1])
	s, _ = strconv.Atoi(arTime[2])
	return
}
func findDigitals(a []int) (rs []int, total int) {
	temp := make(map[int]bool)
	for _, v := range a {
		temp[decimalOne(v)] = true
		temp[decimalTwo(v)] = true
	}
	for k, _ := range temp {
		rs = append(rs, k)
	}
	return rs, len(rs)
}
func findTotalCase(numbers []int, numberToCompare int, powNumber int, isBiggerOrLesser bool) (rs int, next bool) {
	moreLoop := false
	for i := 0; i < len(numbers); i++ {
		switch {
		case numbers[i] == numberToCompare && powNumber == 0:
			rs++
		case numbers[i] == numberToCompare:
			moreLoop = true
		case numbers[i] > numberToCompare && isBiggerOrLesser:
			rs += int(math.Pow(float64(2), float64(powNumber)))
		case numbers[i] < numberToCompare && !isBiggerOrLesser:
			rs += int(math.Pow(float64(2), float64(powNumber)))
		}
	}
	return rs, moreLoop
}
func findTotalCaseReal(hour int, timeCompare Time, bigger bool) (rs int) {
	hourDis, _ := findDigitals([]int{hour})
	numbersCompare := []int{
		decimalOne(timeCompare.minute),
		decimalTwo(timeCompare.minute),
		decimalOne(timeCompare.second),
		decimalTwo(timeCompare.second)}

	if len(hourDis) == 1 {
		for i := 0; i < 6; i++ {
			if hourDis[0] == i {
				rs++
				continue
			}
			temp := []int{hourDis[0], i}
			local := 0
			for pow := 3; pow >= 0; pow-- {
				curRs, next := findTotalCase(temp, numbersCompare[len(numbersCompare)-pow-1], pow, bigger)
				rs += curRs
				local += curRs
				if !next {
					break
				}
			}
		}
	} else {
		for pow := 3; pow >= 0; pow-- {
			curRs, next := findTotalCase(hourDis, numbersCompare[len(numbersCompare)-pow-1], pow, bigger)
			rs += curRs
			if !next {
				break
			}
		}
	}
	return rs
}
func Solution(s string, t string) (rs int) {
	startHo, startMi, startSe := findTimeParts(s)
	endHo, endMi, endSe := findTimeParts(t)
	startTime := Time{
		startHo, startMi, startSe,
	}
	endTime := Time{
		endHo, endMi, endSe,
	}

	totalCaseBiggerStart := findTotalCaseReal(startHo, startTime, true)
	totalCaseLesserEnd := findTotalCaseReal(endHo, endTime, false)
	if startHo == endHo {
		return totalCaseBiggerStart - findTotalCaseReal(startHo, endTime, true)
	}
	if startHo+1 == endHo {
		return totalCaseBiggerStart + totalCaseLesserEnd
	}
	if startHo < endHo {
		first := findTotalCaseReal(startHo, startTime, true)
		mid := findTotalCaseReal(startHo, Time{
			hour:   startHo,
			minute: 00,
			second: 00,
		}, true) * (endHo - startHo - 1)
		last := findTotalCaseReal(endHo, endTime, false)
		rs += first + mid + last
	}
	return rs
}
