package nathan

import "strconv"

type Test struct {
	Name   string
	Result string
}

func isMoreOneTestInGroup(testName string) (rs bool, index int) {
	for i, char := range testName {
		if _, err := strconv.Atoi(string(char)); err == nil {
			return true, i
		}
	}
	return
}

func makeTestList(n []string, r []string) []Test {
	testList := make([]Test, 0)
	for i := 0; i < len(n); i++ {
		testList = append(testList, Test{Name: n[i], Result: r[i]})
	}
	return testList
}

func Solution(t []string, r []string) int {
	testList := makeTestList(t, r)

	groupTest := make(map[string][]Test)
	for _, test := range testList {
		if yes, index := isMoreOneTestInGroup(test.Name); yes {
			list := groupTest[test.Name[:index+1]]
			list = append(list, test)
			groupTest[test.Name[:index+1]] = list
		} else {
			groupTest[test.Name] = []Test{test}
		}
	}

	totalPass := 0
	for _, tests := range groupTest {
		thisGroupPass := true
		for _, test := range tests {
			if test.Result != "OK" {
				thisGroupPass = false
				break
			}
		}
		if thisGroupPass {
			totalPass++
		}
	}
	totalGroup := len(groupTest)

	return totalPass * 100 / totalGroup
}
