package magicsquare

import (
	"math"
)

func findRowColumnNumbers(a [][]int) (row int, column int) {
	row = len(a)
	column = len(a[0])
	return
}

func getSquare(a [][]int, size int, startRow int, startCol int) ([][]int, bool) {
	if startRow > len(a)-size ||
		startCol > len(a[0])-size {
		return make([][]int, 0), true
	}
	rs := make([][]int, 0)
	len := startRow + size
	for i := startRow; i < len; i++ {
		newRow := a[i][startCol : startCol+size]
		rs = append(rs, newRow)
	}
	return rs, false
}

func totalOfS(s []int) (total int) {
	total = 0
	for _, v := range s {
		total += v
	}
	return
}

func isMagicSquare(a [][]int) (int, bool) {
	if len(a) == 0 ||
		len(a[0]) == 0 ||
		len(a) != len(a[0]) {
		return 0, false
	}
	trustTotal := totalOfS(a[0])
	fistDiagonal := make([]int, 0)
	secondDiagonal := make([]int, 0)

	// check all total row
	for i := 0; i < len(a); i++ {
		if totalOfS(a[i]) != trustTotal {
			return 0, false
		}

		// check all total column
		orderCol := make([]int, 0)
		for j := 0; j < len(a); j++ {
			orderCol = append(orderCol, a[j][i])
		}
		if totalOfS(orderCol) != trustTotal {
			return 0, false
		}

		fistDiagonal = append(fistDiagonal, a[i][i])
		secondDiagonal = append(secondDiagonal, a[i][len(a)-i-1])
	}
	// check two diagonals
	if totalOfS(fistDiagonal) != trustTotal ||
		totalOfS(secondDiagonal) != trustTotal {
		return 0, false
	}
	return trustTotal, true
}

func Solution(a [][]int) int {
	row, col := findRowColumnNumbers(a)
	maxSquare := int(math.Max(float64(row), float64(col)))

	for rs := maxSquare; rs > 1; rs-- {

		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a[0]); j++ {
				prepareSquareToCheck, outRange := getSquare(a, rs, i, j)
				if outRange {
					break
				}
				if _, isMagic := isMagicSquare(prepareSquareToCheck); isMagic {
					return rs
				}
			}
		}

	}
	return 1
}
