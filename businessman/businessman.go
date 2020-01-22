package businessman

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getSliceByDateMame(s sort.StringSlice, key string) sort.StringSlice {
	var rs []string
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], key) {
			rs = append(rs, s[i])
		}
	}
	return sort.StringSlice(rs)
}

func findFirstTime(s sort.StringSlice, key string) int {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], key) {
			return i
		}
	}
	return -1
}
func findLatestTime(s sort.StringSlice, key string) int {
	rs := -1
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], key) {
			rs = i
		}
	}
	return rs
}

func findEndTime(s string) string {
	return strings.Split(strings.Split(s, " ")[1], "-")[1]
}
func findStartTime(s string) string {
	return strings.Split(strings.Split(s, " ")[1], "-")[0]
}
func findMaxInt(in []int) int {
	return 0
}
func findDiff(r, s string) int {
	if s == "" {
		s = "24:00"
	}
	if r == "" {
		r = "00:00"
	}
	// r and s : 'hh:mm'
	a := strings.Split(r, ":")
	ahh, _ := strconv.Atoi(a[0])
	amm, _ := strconv.Atoi(a[1])
	b := strings.Split(s, ":")
	bhh, _ := strconv.Atoi(b[0])
	bmm, _ := strconv.Atoi(b[1])

	rs, _ := strconv.Atoi("0")

	if amm > 0 {
		rs += +60 - amm
		ahh++
	}
	sdiff := (bhh - ahh) * 60
	if ahh > bhh {
		return 0
	}
	rs += sdiff
	rs += bmm
	return rs
}

// Solution ...
func Solution(s string) int {
	var (
		Mon = "Mon"
		Tue = "Tue"
		Wed = "Wed"
		Thu = "Thu"
		Fri = "Fri"
		Sat = "Sat"
		Sun = "Sun"
	)

	strAr := strings.Split(s, "\n")
	inputSroted := sort.StringSlice(strAr)
	inputSroted.Sort()

	schedule := []string{
		Mon, Tue, Wed, Thu, Fri, Sat, Sun,
	}

	var freetime []int
	// find all period times
	for _, date := range schedule {
		dateSliceTime := getSliceByDateMame(inputSroted, date)
		for i := 0; i < len(dateSliceTime)-1; i++ {
			freetime = append(freetime, findDiff(findEndTime(dateSliceTime[i]), findStartTime(dateSliceTime[i+1])))
		}
	}
	fmt.Println(freetime)
	return 0
}
