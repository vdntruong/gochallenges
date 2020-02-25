package magicsquare

func findRowColumnNumbers(a [][]int) (row int, column int) {
	row := len(a)
	column := len(a[0])
	return
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getSquare(a [][]int, size int, startRow int, startCol int) [][]int {
	len := startRow + size
	rs := make([][]int, 0)
	for i := startRow; i < len; i++ {
		newRow := append(a[i][startCol:], a[i][:startCol+size]...)
		rs = append(rs, newRow)
	}
	return rs
}

func totalOfS(s []int) (total int) {
	total := 0
	for _, v := range s {
		total += v
	}
	return
}

func isMagicSquare(a [][]int) bool {
	if len(a) != len(a[0]) {
		return false
	}
	trustTotal = totalOfS(a[0])

	// check all total row
	for i := 0; i < len(a); i++ {
		if totalOfS(a[i]) != trustTotal {
			return false
		}
	}

	// check all total column

	return true
}

func Solution(a [][]int) int {
	row, col := findRowColumnNumbers(a)
	maxSquare := findMax(row, col)

	return 0
}
