package reformatnumber

func thisIsNumber(char string) bool {
	return char != " " && char != "-"
}

func Solution(strPhoneNumber string) string {
	var numbers []string
	for _, char := range strPhoneNumber {
		number := string(char)
		if thisIsNumber(number) {
			numbers = append(numbers, number)
		}
	}

	var preRs string
	var subRs string
	preRs = fullOption(numbers)
	if len(numbers)%3 == 1 {
		preRs = fullOption(numbers[:len(numbers)-4])
		subRs = disabilitiesOption(numbers[len(numbers)-4:])
	}
	if preRs == "" {
		return subRs
	}
	if subRs == "" {
		return preRs
	}
	return preRs + "-" + subRs
}

// 123-654
func fullOption(str []string) (rs string) {
	for o := 0; o < len(str); o++ {
		if o != 0 && o%3 == 0 {
			rs += "-"
		}
		rs += str[o]
	}
	return
}

// 12-54
func disabilitiesOption(str []string) (rs string) {
	if len(str) == 4 {
		rs += str[0] + str[1] + "-" + str[2] + str[3]
	}
	return
}
