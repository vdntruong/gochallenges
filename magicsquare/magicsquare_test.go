package magicsquare_test

import (
	"testing"

	"github.com/vdntruong/chl/magicsquare"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc                     string
		inputMatrix              [][]int
		outputMaxSizeMagicSquare int
	}{
		{
			desc: "1 Unit",
			inputMatrix: [][]int{
				[]int{0, 1, 1},
				[]int{1, 1, 1},
				[]int{1, 1, 1},
			},
			outputMaxSizeMagicSquare: 2,
		},
		{
			desc: "1 Unit full size",
			inputMatrix: [][]int{
				[]int{1, 1, 1},
				[]int{1, 1, 1},
				[]int{1, 1, 1},
			},
			outputMaxSizeMagicSquare: 3,
		},
		{
			desc: "Bug in diagonals",
			inputMatrix: [][]int{
				[]int{1, 2, 1},
				[]int{2, 0, 2},
				[]int{1, 2, 1},
			},
			outputMaxSizeMagicSquare: 1,
		},
		{
			desc: "Max size of square matrix",
			inputMatrix: [][]int{
				[]int{4, 9, 2},
				[]int{3, 5, 7},
				[]int{8, 1, 6},
			},
			outputMaxSizeMagicSquare: 3,
		},
		{
			desc: "Max size of square matrix bigger version",
			inputMatrix: [][]int{
				[]int{1, 23, 16, 4, 21},
				[]int{15, 14, 7, 18, 11},
				[]int{24, 17, 13, 9, 2},
				[]int{20, 8, 19, 12, 6},
				[]int{5, 3, 10, 22, 25},
			},
			outputMaxSizeMagicSquare: 5,
		},
		{
			desc: "Max size of square matrix other bigger version",
			inputMatrix: [][]int{
				[]int{4, 9, 2, 4, 9, 2},
				[]int{3, 5, 7, 3, 5, 7},
				[]int{8, 1, 6, 8, 1, 6},
				[]int{4, 9, 2, 4, 9, 2},
				[]int{3, 5, 7, 3, 5, 7},
				[]int{8, 1, 6, 8, 1, 6},
			},
			outputMaxSizeMagicSquare: 6,
		},
		{
			desc: "Non magic square, max size is ONE",
			inputMatrix: [][]int{
				[]int{0, 5, 2, 5, 3},
				[]int{3, 1, 0, 8, 4},
				[]int{0, 0, 0, 2, 0},
				[]int{8, 4, 9, 0, 0},
			},
			outputMaxSizeMagicSquare: 1,
		},
		{
			desc: "Total zero",
			inputMatrix: [][]int{
				[]int{4, 0, 4, 5, 3},
				[]int{2, 7, 3, 8, 4},
				[]int{1, 7, 6, 0, 0},
				[]int{8, 4, 9, 0, 0},
			},
			outputMaxSizeMagicSquare: 2,
		},
		{
			desc: "Test case 01 of challenge",
			inputMatrix: [][]int{
				[]int{4, 3, 4, 5, 3},
				[]int{2, 7, 3, 8, 4},
				[]int{1, 7, 6, 5, 2},
				[]int{8, 4, 9, 5, 5},
			},
			outputMaxSizeMagicSquare: 3,
		},
		{
			desc: "Test case 02 of challenge",
			inputMatrix: [][]int{
				[]int{2, 2, 1, 1},
				[]int{2, 2, 2, 2},
				[]int{1, 2, 2, 2},
			},
			outputMaxSizeMagicSquare: 2,
		},
		{
			desc: "Test case 03 of challenge",
			inputMatrix: [][]int{
				[]int{7, 2, 4},
				[]int{2, 7, 6},
				[]int{9, 5, 1},
				[]int{4, 3, 8},
				[]int{3, 5, 4},
			},
			outputMaxSizeMagicSquare: 3,
		},
		{
			desc: "Includes multiple result (3,2)",
			inputMatrix: [][]int{
				[]int{0, 0, 2, 4, 9, 2},
				[]int{0, 0, 7, 3, 5, 7},
				[]int{8, 1, 6, 8, 1, 6},
				[]int{4, 9, 2, 4, 9, 2},
				[]int{3, 5, 7, 3, 5, 7},
				[]int{8, 1, 6, 8, 1, 3},
			},
			outputMaxSizeMagicSquare: 3,
		},
		{
			desc: "Big matrix with big magic square",
			inputMatrix: [][]int{
				[]int{6, 4, 9, 2, 4, 9, 2, 6},
				[]int{2, 4, 9, 2, 4, 9, 2, 4},
				[]int{8, 3, 5, 7, 3, 5, 7, 7},
				[]int{0, 8, 1, 6, 8, 1, 6, 0},
				[]int{7, 4, 9, 2, 4, 9, 2, 1},
				[]int{6, 3, 5, 7, 3, 5, 7, 6},
				[]int{3, 8, 1, 6, 8, 1, 6, 8},
				[]int{3, 8, 5, 7, 8, 3, 6, 3},
			},
			outputMaxSizeMagicSquare: 6,
		},
		{
			desc: "Big matrix with multiple magic squares: 7,5,3",
			inputMatrix: [][]int{
				[]int{46, 8, 16, 20, 29, 7, 49},
				[]int{3, 40, 35, 36, 18, 41, 2},
				[]int{44, 12, 33, 23, 19, 38, 6},
				[]int{28, 26, 11, 25, 39, 24, 22},
				[]int{5, 37, 31, 27, 17, 13, 45},
				[]int{48, 9, 15, 14, 32, 10, 47},
				[]int{1, 43, 34, 30, 21, 42, 4},
			},
			outputMaxSizeMagicSquare: 7,
		},
		{
			desc: "Big matrix with multiple magic squares: 3",
			inputMatrix: [][]int{
				[]int{46, 8, 16, 10, 29, 7, 49},
				[]int{3, 40, 35, 36, 18, 41, 2},
				[]int{44, 12, 33, 23, 19, 38, 6},
				[]int{28, 26, 11, 25, 39, 24, 22},
				[]int{5, 37, 31, 27, 17, 13, 45},
				[]int{48, 8, 15, 14, 32, 10, 47},
				[]int{1, 43, 34, 30, 21, 42, 4},
			},
			outputMaxSizeMagicSquare: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := magicsquare.Solution(tC.inputMatrix); rs != tC.outputMaxSizeMagicSquare {
				t.Errorf("Got max size=%d, wants size=%d", rs, tC.outputMaxSizeMagicSquare)
			}
		})
	}
}
