package lovelynumber

import (
	"strconv"
	"strings"
)

// a and b in range [0..100000]
// a <= b
func Solution(a, b int) int {
	rs := 0
	for i := a; i <= b; i++ {
		if !notLovely(i) {
			rs++
		}
	}
	return rs
}

func notLovely(i int) (rs bool) {
	if i < 111 {
		return
	}
	iStr := strconv.FormatInt(int64(i), 10)
	for i := 0; i < len(iStr); i++ {
		if strings.Count(iStr, string(iStr[i])) >= 3 {
			return true
		}
	}
	return
}
