package cuttingtree

import "github.com/sirupsen/logrus"

func stillOk(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func Solution(a []int) (rs int) {
	for i := 0; i < len(a)-1; i++ {
		ar := append(a[:i], a[i+1:]...)
		logrus.Info(a)
		if stillOk(ar) {
			rs++
		}
	}
	logrus.Info(a)
	// if stillOk(a[:len(a)-1]) {
	// 	rs++
	// }
	return rs
}
