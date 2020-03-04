package johnvacation

var (
	months = []string{
		"January", "February", "March", "April",
		"May", "June", "July", "August", "September",
		"October", "November", "December",
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
	numberDayInMonths = map[string]int{
		"January": 31, "February": 28, "March": 31,
		"April": 30, "May": 31, "June": 30,
		"July": 31, "August": 31, "September": 30,
		"October": 31, "November": 30, "December": 31,
	}
)

func findTotalDayOfMonth(y int, m string) (rs int) {
	rs += numberDayInMonths[m]
	if monthsOfYear[m] == 2 && y%4 == 0 {
		rs++
	}
	return rs
}
func findTotalDayBeforeMonth(y int, m string) (rs int) {
	for _, month := range months {
		if month != m {
			rs += findTotalDayOfMonth(y, m)
		} else {
			break
		}
	}
	return rs
}
func findTotalDayAfterMonth(y int, m string) int {
	return findTotalDayBeforeMonth(y, m) + findTotalDayOfMonth(y, m)
}
func findIndexOfMonth(m string) int {
	for i := 0; i < len(months); i++ {
		if months[i] == m {
			return i
		}
	}
	return -1
}

func getFirstDayOfWeekInMonth(y int, m string, firstWeekdayInYear string) (rs string) {
	totalDays := findTotalDayBeforeMonth(y, m)
	extraDays := totalDays % 7
	for i, dayOfWeek := range days {
		if dayOfWeek == firstWeekdayInYear {
			dayIndex := i + extraDays
			if dayIndex > len(days)-1 {
				dayIndex = dayIndex%(len(days)-1) - 1
			}
			rs = days[dayIndex]
			break
		}
	}
	return rs
}
func getFirstMondayDay(y int, m string, firstWeekdayInYear string) (rs int) {
	firstDayOfMonth := getFirstDayOfWeekInMonth(y, m, firstWeekdayInYear)
	if firstDayOfMonth == days[0] {
		return 1
	}
	extraDays := 0
	for i, day := range days {
		if day == firstDayOfMonth {
			extraDays = len(days) + 1 - i
			break
		}
	}
	return rs + extraDays
}
func getLastSundayDay(y int, m string, firstWeekdayInYear string) (rs int) {
	dayOfFirstMonday := getFirstMondayDay(y, m, firstWeekdayInYear)
	totalDaysInMonth := numberDayInMonths[m]
	if y%4 == 0 && totalDaysInMonth == 28 {
		totalDaysInMonth++
	}
	return dayOfFirstMonday + ((totalDaysInMonth-dayOfFirstMonday)/7)*7 - 1
}
func getMonthMiddle(s, e string) (rs []string) {
	for i := 0; i < len(months); i++ {
		if i > findIndexOfMonth(s) && i < findIndexOfMonth(e) {
			rs = append(rs, months[i])
		}
	}
	return rs
}

func Solution(year int, startMonth, endMonth, weekday string) int {
	totalPreDays := findTotalDayOfMonth(year, startMonth) - getFirstMondayDay(year, startMonth, weekday) + 1
	totalPostDays := getLastSundayDay(year, endMonth, weekday)
	totalMiddleDays := 0
	listMiddleMonths := getMonthMiddle(startMonth, endMonth)
	if len(listMiddleMonths) > 0 {
		for _, month := range listMiddleMonths {
			totalMiddleDays += findTotalDayOfMonth(year, month)
		}
	}
	return (totalPreDays + totalMiddleDays + totalPostDays) / 7
}
