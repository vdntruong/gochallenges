package sortingarray

func arNotOke(a []int) (rs bool) {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return true
		}
	}
	return
}

func Solution(a []int) bool {
	if len(a) < 3 {
		return true
	}

	evilIndex := -1
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			evilIndex = i
			break
		}
	}
	if evilIndex == -1 {
		return true
	}
	indexToChange := -1
	for i := 1; i < len(a); i++ {
		if i != evilIndex && a[i] >= a[evilIndex] {
			indexToChange = i - 1
			break
		}
	}
	if indexToChange == -1 {
		return false
	}
	temp := a[evilIndex]
	a[evilIndex] = a[indexToChange]
	a[indexToChange] = temp
	return !arNotOke(a)
}
