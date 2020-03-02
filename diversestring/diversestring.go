package diversestring

func stopNow(a, b, c int) bool {
	return (a == 0 && b == 0) ||
		(a == 0 && c == 0) ||
		(b == 0 && c == 0)
}

func addStr(str, s string, times int) string {
	for i := 0; i < times; i++ {
		str += s
	}
	return str
}

func makeRs(disCenter map[string]int) (rs string) {
	for {
		if stopNow(disCenter["a"], disCenter["b"], disCenter["c"]) {
			return rs
		}
		rs += addStr(rs, "a", disCenter["a"])
	}
}

func Solution(a, b, c int) (rs string) {
	distributeCenter := map[string]int{"a": a, "b": b, "c": c}
	rs = makeRs(distributeCenter)
	return rs
}
