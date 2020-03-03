package interestingtime

import (
	"strconv"
	"strings"
)

type (
	Tim struct {
		hour, minute, second int
	}
)

func makeSense(str string) (hour, minute, second int) {
	arTime := strings.Split(str, ":")
	hour, _ = strconv.Atoi(arTime[0])
	minute, _ = strconv.Atoi(arTime[1])
	second, _ = strconv.Atoi(arTime[2])
	return
}
func countDigitals(hour, minute, second int) (rs []int, total int) {
	temp := make(map[int]bool)
	temp[hour/10] = true
	temp[hour%10] = true
	temp[minute/10] = true
	temp[minute%10] = true
	temp[second/10] = true
	temp[second%10] = true
	for k, _ := range temp {
		rs = append(rs, k)
	}
	return rs, len(rs)
}
func totalInStand(a []int, stand int) (rs int) {
	for _, v := range a {
		if v <= stand {
			rs++
		}
	}
	return rs
}
func makeTimes(a []int) (rs []Tim) {
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 2; j++ {

	// 	}
	// }
	return rs
}

func Solution(s string, t string) (rs int) {
	// startHour, _, _ := makeSense(s)
	// endHour, _, _ := makeSense(t)
	// listRs := []Tim{}

	// ls, c := countDigitals(startHour, startHour, startHour)
	// startIndex := startHour
	// for startIndex <= endHour {

	// }
	// return len(listRs)
	return
}
