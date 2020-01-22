package main_test

import (
	"fmt"
	"strings"
	"testing"
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
			return canBeating(B, curRow-2, curCol-2) + 1 + canBeating(B, curRow-2, curCol+2) + 1
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
func solution(B []string) int {
	curRow := 0
	curCol := 0
	for i := 0; i < len(B); i++ {
		if strings.Contains(B[i], `O`) {
			curRow = i
			for j := 0; j < len(B[i]); j++ {
				if string(B[i][j]) == `O` {
					curCol = j
				}
			}
		}
	}
	fmt.Printf("Index of 0(%v:%v) and rs is: ", curRow, curCol)
	return canBeating(B, curRow, curCol)
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		n string
		i []string
		o int
	}{
		{
			n: "One",
			i: []string{
				`..X...`,
				`......`,
				`....X.`,
				`.X....`,
				`..X.X.`,
				`...O..`},
			o: 2,
		}, {
			n: "Two",
			i: []string{
				`X....`,
				`.X...`,
				`..O..`,
				`...X.`,
				`.....`},
			o: 0,
		},
	}

	for _, test := range testCases {
		t.Run(test.n, func(t *testing.T) {
			got := solution(test.i)
			if got != test.o {
				t.Errorf("got %v, wants %v", got, test.o)
			}
		})
	}
}
