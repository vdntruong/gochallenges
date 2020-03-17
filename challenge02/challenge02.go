package challenge02

import (
	"fmt"
	"regexp"
)

func Solution(a, b, c, d int) (rs int) {
	cases := deduplicate(shuffle([]int{a, b, c, d}))
	for _, c := range cases {
		if validateTime(c) {
			rs++
		}
	}
	return rs
}

var timeRegexp = regexp.MustCompile("^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$")

func validateTime(t string) bool {
	return timeRegexp.MatchString(t)
}

func deduplicate(arr []string) []string {
	m := make(map[string]bool)
	for _, s := range arr {
		m[s] = true
	}

	result := []string{}
	for k := range m {
		result = append(result, k)
	}
	return result
}

func shuffle(arr []int, nums ...int) []string {
	result := []string{}
	if len(arr)+len(nums) < 4 {
		return []string{}
	}

	if len(nums) == 3 {
		return gen(nums[0], nums[1], nums[2], arr)
	}

	for i, num := range arr {
		newNums := []int{}
		newNums = append(newNums, nums...)
		newNums = append(newNums, num)

		newArr := []int{}
		newArr = append(newArr, arr[:i]...)
		newArr = append(newArr, arr[i+1:]...)

		result = append(result, shuffle(newArr, newNums...)...)
	}
	return result
}

func gen(a, b, c int, ds []int) []string {
	result := []string{}
	for _, d := range ds {
		result = append(result, fmt.Sprintf("%d%d:%d%d", a, b, c, d))
	}
	return result
}
