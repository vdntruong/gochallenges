package johnvacation

var (
	months = []string{
		"January", "February", "March", "April",
		"May", "June", "July", "August", "September",
		"October", "November", "December",
	}
	numberDayInMonths = map[string]int{
		"January": 31, "February": 28, "March": 31,
		"April": 30, "May": 31, "June": 30,
		"July": 31, "August": 31, "September": 30,
		"October": 31, "November": 30, "December": 31,
	}
	days = []string{
		"Monday", "Tuesday", "Wednesday",
		"Thursday", "Friday", "Saturday", "Sunday",
	}
	monthsOfYear = map[string]int{
		"January": 1, "February": 2, "March": 3,
		"April": 4, "May": 5, "June": 6,
		"July": 7, "August": 8, "September": 9,
		"October": 10, "November": 11, "December": 12,
	}
)

func totalDayOfMonth(y int, m string) (rs int) {
	rs += numberDayInMonths[m]
	if m == months[1] && y%4 == 0 {
		rs++
	}
	return rs
}
func findTotalDayBeforeMonth(y int, m string) (rs int) {
	for i := 0; i < len(months); i++ {
		if months[i] == m {
			break
		}
		rs += totalDayOfMonth(y, months[i])
	}
	return rs
}
func findTotalDayToMonth(y int, m string) int {
	return findTotalDayBeforeMonth(y, m) + totalDayOfMonth(y, m)
}
func findIndexOfMonth(m string) int {
	for i := 0; i < len(months); i++ {
		if months[i] == m {
			return i
		}
	}
	return -1
}

func getFirstDayOfWeekInMonth(y int, m string, firstWeekdayInYear string) string {
	totalDays := findTotalDayBeforeMonth(y, m)
	extraDays := totalDays % 7
	firstDayOfYearIndex := 0
	for i, dayOfWeek := range days {
		if dayOfWeek == firstWeekdayInYear {
			firstDayOfYearIndex = i
			break
		}
	}
	if firstDayOfYearIndex+extraDays >= len(days) {
		return days[(firstDayOfYearIndex+extraDays)%7]
	}
	return days[firstDayOfYearIndex+extraDays]
}
func getFirstMondayDay(y int, m string, firstWeekdayInYear string) (rs int) {
	firstDayOfMonth := getFirstDayOfWeekInMonth(y, m, firstWeekdayInYear)
	if firstDayOfMonth == days[0] {
		return 1
	}
	for i, day := range days {
		if day == firstDayOfMonth {
			return len(days) + 1 - i
		}
	}
	return
}
func getLastSundayDay(y int, m string, firstWeekdayInYear string) (rs int) {
	dayOfFirstMonday := getFirstMondayDay(y, m, firstWeekdayInYear)
	totalDaysInMonth := totalDayOfMonth(y, m)
	return ((totalDaysInMonth-dayOfFirstMonday+1)/7)*7 + dayOfFirstMonday - 1
}

func getMonthMiddle(s, e string) (rs []string) {
	for i := 0; i < len(months); i++ {
		if findIndexOfMonth(s) < i && i < findIndexOfMonth(e) {
			rs = append(rs, months[i])
		}
	}
	return rs
}

func Solution(year int, startMonth, endMonth, weekday string) int {
	if startMonth == endMonth {
		firstMonday := getFirstMondayDay(year, startMonth, weekday)
		lastSunday := getLastSundayDay(year, endMonth, weekday)
		return (lastSunday - firstMonday + 1) / 7
	}
	totalPreDays := totalDayOfMonth(year, startMonth) - getFirstMondayDay(year, startMonth, weekday) + 1
	totalPostDays := getLastSundayDay(year, endMonth, weekday)
	totalMiddleDays := 0
	listMiddleMonths := getMonthMiddle(startMonth, endMonth)
	if len(listMiddleMonths) > 0 {
		for _, month := range listMiddleMonths {
			totalMiddleDays += totalDayOfMonth(year, month)
		}
	}
	return (totalPreDays + totalMiddleDays + totalPostDays) / 7
}

// Testing
func FistMonday(y int, m string, fwd string) int {
	return getFirstMondayDay(y, m, fwd)
}
func LastSunday(y int, m string, fwd string) int {
	return getLastSundayDay(y, m, fwd)
}
