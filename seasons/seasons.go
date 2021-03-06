package seasons

import (
	"sort"
)

func findMaxAmplitude(ar []int) int {
	sort.Ints(ar)
	return ar[len(ar)-1] - ar[0]
}

func Solution(t []int) (rs string) {
	daysPerSeason := len(t) / 4
	startTempIndex := 0
	maxAmplitude := -1
	for _, season := range []string{"WINTER", "SPRING", "SUMMER", "AUTUMN"} {
		maxTempeOfSeason := findMaxAmplitude(t[startTempIndex : startTempIndex+daysPerSeason])
		if maxTempeOfSeason > maxAmplitude {
			maxAmplitude = maxTempeOfSeason
			rs = season
		}
		startTempIndex += daysPerSeason
	}
	return rs
}
