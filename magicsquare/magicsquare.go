package magicsquare

import "math"

func findRowColumnNumbers(a [][]int) (row int, column int) {
	row = len(a)
	column = len(a[0])
	return
}

func getSquare(a [][]int, size int, startRow int, startCol int) [][]int {
	if startRow + size - 1 > len(a){
		return make([][]int, 0)
	}
	if startCol + size - 1 > len(a[0]){
		return make([][]int, 0)
	}
	if startRow > len(a) - size {
		return make([][]int, 0)
	}
	if startCol > len(a[0]) - size {
		return make([][]int, 0)
	}
	len := startRow + size
	rs := make([][]int, 0)
	for i := startRow; i < len; i++ {
		newRow := a[i][startCol:startCol+size]
		rs = append(rs, newRow)
	}
	return rs
}

func totalOfS(s []int) (total int) {
	total = 0
	for _, v := range s {
		total += v
	}
	return
}

func isMagicSquare(a [][]int) bool {
	if len(a) == 0 || len(a[0])==0{
		return false
	}
	if len(a) != len(a[0]) {
		return false
	}
	trustTotal := totalOfS(a[0])
	fistDiagonal := make([]int, 0)
	secondDiagonal := make([]int, 0)

	// check all total row
	for i := 0; i < len(a); i++ {
		if totalOfS(a[i]) != trustTotal {
			return false
		}

		// check all total column
		orderCol := make([]int, 0)
		for j:=0; j<len(a);j++ {
			orderCol = append(orderCol, a[j][i])
		}
		if totalOfS(orderCol) != trustTotal{
			return false
		}

		fistDiagonal = append(fistDiagonal, a[i][i])
		secondDiagonal = append(secondDiagonal, a[i][len(a) - i])
	}
	// checl two diagonals
	if totalOfS(fistDiagonal) != trustTotal {
		return false
	}
	if totalOfS(secondDiagonal) != trustTotal {
		return false
	}
	return true
}

func Solution(a [][]int) int {
	row, col := findRowColumnNumbers(a)
	maxSquare := int(math.Max(float64(row), float64(col)))

	for rs:= maxSquare; rs >1; rs++ {
		for i:= 0; i<len(a); i++{
			for j:= 0; j<len(a[0]); j++{
				prepareSquareToCheck := getSquare(a, rs, i, j)
				if isMagicSquare(prepareSquareToCheck){
					return rs
				}
			}
		}
	}
	

	return 1
}
