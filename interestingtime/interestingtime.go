package interestingtime

import (
	"strconv"
	"strings"
)

func Solution(s string, t string) (rs int) {
	startHour, startMinute, startSecond := makeSence(s)
	endHour, endMinute, endSecond := makeSence(t)

	for h := startHour; h<=endHour; h++{
		for m:= startMinute; m<=endMinute; m++{
			for s:= startSecond; s<=endSecond; s++{
				if _, count := countDigitals(h, m, s); count <= 2{
					rs++
				}
			}
		}
	}
	
	return rs
}
func makeSence(str string) (hour, minute, second int) {
	arTime := strings.Split(str, ":")
	hour, _ = strconv.Atoi(arTime[0])
	minute, _ = strconv.Atoi(arTime[1])
	second, _ = strconv.Atoi(arTime[2])
	return
}
func countDigitals(hour, minute, second int) (rs []int, total int) {
	temp := make(map[int]bool)
	temp[hour/10] = true
	temp[minute/10] = true
	temp[second/10] = true
	temp[hour%10] = true
	temp[minute%10] = true
	temp[second%10] = true
	for k, _ := range temp {
		rs = append(rs, k)
	}
	return rs, len(rs)
}
