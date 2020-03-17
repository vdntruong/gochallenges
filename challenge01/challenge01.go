package challenge01

import "unicode"

var (
	CAPSLOCK  = 'C'
	BACKSPACE = 'B'
)

func Solution(s string) string {
	rs := []rune{}
	isCapslock := false

	for _, c := range s {
		if c == CAPSLOCK {
			isCapslock = !isCapslock
			continue
		}

		if c == BACKSPACE {
			if len(rs) != 0 {
				rs = rs[:len(rs)-1]
			}
			continue
		}

		if isCapslock {
			rs = append(rs, unicode.ToUpper(c))
		} else {
			rs = append(rs, c)
		}
	}

	return string(rs)
}
