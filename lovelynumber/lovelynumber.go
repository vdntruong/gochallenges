package lovelynumber

import (
	"strconv"
	"strings"
)

// a and b in range [0..100000]
// a <= b
func Solution(a, b int) int {
	if a > b {
		return 0
	}

	rs := 0
	for i := a; i <= b; i++ {
		if !notLovely(i) {
			rs++
		}
	}

	return rs
}

func notLovely(i int) bool {
	if i < 111 {
		return false
	}

	istr := strconv.FormatInt(int64(i), 10)
	for i, _ := range istr {
		if strings.Count(istr, string(istr[i])) >= 3 {
			return true
		}
	}

	return false
}
