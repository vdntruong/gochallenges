package equilibriummatrix

func findRowCol(a [][]int) (row, col int){
	return len(a), len(a[0])
}
func findWeight(a[]int) (rs int){
	for _, v := range a{
		rs+=v
	}
	return rs
}
func findAllWeight(a [][]int) (map[int]int, map[int]int) {
	weightRow := make(map[int]int)
	weightCol := make(map[int]int)

	row, col := findRowCol(a)
	doneRow := false
	for i:= 0; i<col; i++{
		weCol := 0
		for j:= 0; j<row; j++{
			if !doneRow {
				weightRow[j] = findWeight(a[j])
			}
			if j == row-1{
				doneRow = true
			}
			weCol += a[j][i]
		}
		weightCol[i] = weCol
	}
	return weightRow, weightCol
}
func okUpDownRow(i, row int, weightRow map[int]int) bool {
	upRow := 0
	for x:=i-1; x>=0; x--{
		upRow += weightRow[x]
	}
	downRow := 0
	for x:=i+1; x<row; x++{
		downRow += weightRow[x]
	}
	return upRow == downRow
}
func okLeftRightCol(j, col int, weightCol map[int]int) bool {
	leftCol := 0
	for x:=j-1; x>=0; x--{
		leftCol += weightCol[x]
	}
	rightCol := 0
	for x:=j+1; x<col; x++{
		rightCol += weightCol[x]
	}
	return leftCol == rightCol
}
func isEquilibriumPoint(i, j, row, col int, weightRow, weightCol map[int]int) bool {
	return okUpDownRow(i, row, weightRow) && okLeftRightCol(j, col, weightCol)
}

func Solution(a [][]int) (rs int) {
	row, col := findRowCol(a)
	weightRow, weightCol := findAllWeight(a)
	
	for i:= 0; i<row-1; i++{
		for j:= 0; j<col-1; j++{
			if isEquilibriumPoint(i, j, row, col, weightRow, weightCol){
				rs++
			}
		}
	}
	return rs
}
