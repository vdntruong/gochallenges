package rotatedice

import "github.com/sirupsen/logrus"

func contain(a []int, v int) bool {
	for _, x := range a {
		if x == v {
			return true
		}
	}
	return false
}

func countVisibleTime(a []int, v int) (rs int) {
	for _, x := range a {
		if x == v {
			rs++
		}
	}
	return
}

func Solution(a []int) int {
	oppositeFaces := map[int]int{
		1: 6, 6: 1,
		2: 5, 5: 2,
		3: 4, 4: 3,
	}
	logrus.Info(oppositeFaces)
	countFaceList := make(map[int]int)
	maxCount := 0
	maxFaceCount := -1
	for _, face := range a {
		if countFaceList[face] == 0 {
			counted := countVisibleTime(a, face)
			countFaceList[face] = counted
			if counted > maxCount {
				maxCount = counted
				maxFaceCount = face
			}
		}
	}

	if maxFaceCount == -1 {

	}

	return 0
}
