package jafarcheck

import (
	"strings"
)

func isX(B []string, i, j int) bool {
	return string(B[i][j]) == "X"
}
func canBeating(B []string, curRow, curCol int) int {
	if curRow <= 1 {
		return 0
	}
	if curCol > 1 && curCol < len(B)-2 {
		if isX(B, curRow-1, curCol-1) && !isX(B, curRow-2, curCol-2) && isX(B, curRow-1, curCol+1) && !isX(B, curRow-2, curCol+2) {
			if canBeating(B, curRow-2, curCol-2) > canBeating(B, curRow-2, curCol+2) {
				return canBeating(B, curRow-2, curCol-2) + 1
			}
			return canBeating(B, curRow-2, curCol+2) + 1
		}
	}
	if curCol > 1 {
		if isX(B, curRow-1, curCol-1) && !isX(B, curRow-2, curCol-2) {
			return canBeating(B, curRow-2, curCol-2) + 1
		}
	}
	if curCol < len(B)-2 {
		if isX(B, curRow-1, curCol+1) && !isX(B, curRow-2, curCol+2) {
			return canBeating(B, curRow-2, curCol+2) + 1
		}
	}
	return 0
}
func Solution(B []string) int {
	curRow := 0
	curCol := 0
	for i := 0; i < len(B); i++ {
		if strings.Contains(B[i], `O`) {
			curRow = i
			for j := 0; j < len(B[i]); j++ {
				if string(B[i][j]) == `O` {
					curCol = j
					break
				}
			}
			break
		}
	}
	return canBeating(B, curRow, curCol)
}
