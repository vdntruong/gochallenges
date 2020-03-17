package main

import (
	"unicode"
)

var (
	BACKSPACE = 'B'
	CAPSLOCK  = 'C'
)

func Solution(s string) string {
	isActiveCL := false
	rs := []rune{}
	for _, r := range s {
		if r == CAPSLOCK {
			isActiveCL = !isActiveCL
			continue
		}

		if r == BACKSPACE {
			if len(rs) != 0 {
				rs = rs[:len(rs)-1]
			}
			continue
		}

		if isActiveCL {
			rs = append(rs, unicode.ToUpper(r))
		} else {
			rs = append(rs, r)
		}
	}
	return string(rs)
}
